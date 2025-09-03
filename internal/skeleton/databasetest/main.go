package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"github.com/yuita-yoshihiko/daredemo-design-backend/internal/skeleton"
)

func main() {
	inRoot := "usecase/repository"
	outDir := "adapter/database"

	names := collectBasenames(inRoot)
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		log.Fatal(err)
	}

	const headerTpl = `package database_test

import (
	"testing"
	"time"

	"github.com/yuita-yoshihiko/daredemo-design-backend/adapter/database"
	"github.com/yuita-yoshihiko/daredemo-design-backend/infrastructure/db"
	"github.com/yuita-yoshihiko/daredemo-design-backend/models"
	"github.com/yuita-yoshihiko/daredemo-design-backend/testutils"
)
`

	const testTpl = `
func Test_{{ .Upper }}_{{ .Method }}(t *testing.T) {
	data := testutils.LoadFixture(t, "testfixtures/{{ .Snake }}s/{{ .MethodSnake }}")
	dbUtils := db.NewDBUtil(data)
	r := database.New{{ .Upper }}Repository(dbUtils)

	type args struct {
		id int64
	}

	tests := []struct {
		name string
		args args
		want *models.{{ .Upper }}
	}{
		{
			name: "idで抽出した単一の{{ .Upper }}のデータが取得できる",
			args: args{ id: 1 },
			want: &models.{{ .Upper }}{
				// TODO: 必要なフィールドを追加してください
				ID:        1,
				CreatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.{{ .Method }}(t.Context(), tt.args.id)
			if err != nil {
				t.Errorf("error = %v", err)
			}
			testutils.AssertResponse(t, got, tt.want, models.{{ .Upper }}{})
		})
	}
}
`
	testT := template.Must(template.New("test").Parse(testTpl))

	// 「(ctx context.Context」の直前の識別子＝メソッド名 を抽出
	methodRe := regexp.MustCompile(`([A-Za-z_][A-Za-z0-9_]*)\s*\(\s*context\.Context\b`)
	for name := range names {
		srcPath := filepath.Join("usecase", "repository", name+".go")
		b, err := os.ReadFile(srcPath)
		if err != nil {
			log.Fatalf("read %s: %v", srcPath, err)
		}

		raw := methodRe.FindAllStringSubmatch(string(b), -1)
		seen := map[string]struct{}{}
		var methods []string
		for _, m := range raw {
			if len(m) < 2 {
				continue
			}
			fn := m[1]
			if _, ok := seen[fn]; ok {
				continue
			}
			seen[fn] = struct{}{}
			methods = append(methods, fn)
		}
		if len(methods) == 0 {
			fmt.Printf("no methods found in %s (skip)\n", srcPath)
			continue
		}
		sort.Strings(methods) // 安定化

		// test file path
		outPath := filepath.Join(outDir, name+"_test.go")

		if _, err := os.Stat(outPath); os.IsNotExist(err) {
			// 新規作成: ヘッダー + 各メソッドのテスト
			var buf bytes.Buffer
			buf.WriteString(headerTpl)

			for _, mname := range methods {
				data := struct {
					skeleton.Name
					Method      string
					MethodSnake string
				}{
					Name:        skeleton.Name(name), // {{ .Upper }} / {{ .Snake }}
					Method:      mname,               // Fetch, Create, ...
					MethodSnake: strcase.ToSnake(mname),
				}

				// フィクスチャも作成
				if err := ensureFixture(data.MethodSnake, data.Name.Snake()); err != nil {
					log.Fatalf("ensure fixture failed: %v", err)
				}

				if err := testT.Execute(&buf, data); err != nil {
					log.Fatalf("tpl execute: %v", err)
				}
			}

			src := buf.Bytes()
			if fmtd, err := format.Source(src); err == nil {
				src = fmtd
			}
			if err := os.WriteFile(outPath, src, 0o644); err != nil {
				log.Fatalf("write %s: %v", outPath, err)
			}
			fmt.Printf("created: %s\n", outPath)
			continue
		}

		// 既存テスト: 無いメソッドだけ追記 + フィクスチャ生成
		appendMissing(outPath, skeleton.Name(name), methods, testT)
		for _, mname := range methods {
			if err := ensureFixture(strcase.ToSnake(mname), skeleton.Name(name).Snake()); err != nil {
				log.Fatalf("ensure fixture failed: %v", err)
			}
		}
	}
}

// 既存 test ファイルに無いテスト関数だけ追記
func appendMissing(outPath string, n skeleton.Name, methods []string, t *template.Template) {
	existing, _ := os.ReadFile(outPath)
	exists := string(existing)

	var add bytes.Buffer
	for _, mname := range methods {
		funcSig := fmt.Sprintf("func Test_%s_%s(", n.Upper(), mname)
		if strings.Contains(exists, funcSig) {
			continue
		}
		data := struct {
			skeleton.Name
			Method      string
			MethodSnake string
		}{
			Name:        n,
			Method:      mname,
			MethodSnake: strcase.ToSnake(mname),
		}
		if err := t.Execute(&add, data); err != nil {
			log.Fatalf("tpl execute(append): %v", err)
		}
	}

	if add.Len() == 0 {
		fmt.Printf("skip (no new methods): %s\n", outPath)
		return
	}

	seg := add.Bytes()
	// 可能なら全体フォーマットして上書き
	if fmtd, err := format.Source(append(existing, seg...)); err == nil {
		if err := os.WriteFile(outPath, fmtd, 0o644); err != nil {
			log.Fatalf("rewrite %s: %v", outPath, err)
		}
	} else {
		f, err := os.OpenFile(outPath, os.O_WRONLY|os.O_APPEND, 0o644)
		if err != nil {
			log.Fatalf("open %s: %v", outPath, err)
		}
		defer f.Close()
		if _, err := f.Write(append([]byte("\n"), seg...)); err != nil {
			log.Fatalf("append %s: %v", outPath, err)
		}
	}
	fmt.Printf("appended methods to: %s\n", outPath)
}

func ensureFixture(methodSnake, entitySnake string) error {
	dir := filepath.Join("adapter", "database", "testfixtures", pluralizer(entitySnake), methodSnake)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	file := filepath.Join(dir, pluralizer(entitySnake)+".yml")
	if _, err := os.Stat(file); err == nil {
		// 既にある
		return nil
	} else if !os.IsNotExist(err) {
		return err
	}

	content := fmt.Sprintf("- id: 1" + "\n")

	return os.WriteFile(file, []byte(content), 0o644)
}

// usecase/repository配下の .go ベース名を一意に収集（再帰）
func collectBasenames(root string) map[string]struct{} {
	out := make(map[string]struct{})
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if strings.ToLower(filepath.Ext(d.Name())) != ".go" {
			return nil
		}
		if strings.HasSuffix(d.Name(), "_test.go") {
			return nil
		}
		base := strings.TrimSuffix(d.Name(), filepath.Ext(d.Name()))
		out[base] = struct{}{}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func pluralizer(s string) string {
	plu := pluralize.NewClient()
	return plu.Plural(s)
}
