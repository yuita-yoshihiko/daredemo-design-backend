
-- +migrate Up
CREATE TABLE IF NOT EXISTS design_tip_categories (
  design_tip_id BIGINT NOT NULL REFERENCES design_tips(id),
  category_id BIGINT NOT NULL REFERENCES categories(id),
  PRIMARY KEY (design_tip_id, category_id)
);

comment on table design_tip_categories is 'デザイン情報とカテゴリーの中間テーブル';

-- +migrate Down
DROP TABLE IF EXISTS design_tip_categories;
