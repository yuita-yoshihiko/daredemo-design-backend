

-- +migrate Up
CREATE TABLE IF NOT EXISTS design_tips (
  id BIGSERIAL,
  title TEXT NOT NULL,
  guidance TEXT NOT NULL,
  url TEXT NOT NULL,
  media TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id)
);

comment on table design_tips is 'デザイン情報';
comment on column design_tips.title is 'タイトル';
comment on column design_tips.guidance is 'ガイダンス';
comment on column design_tips.url is 'URL';
comment on column design_tips.media is '媒体';

-- +migrate Down
DROP TABLE IF EXISTS design_tips;
