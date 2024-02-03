# rssagg
Free time GO lang project to build a rss aggregrator

# goose
up migration:
goose postgres postgres://<userName>:<password>@localhost:5432/rssagg up
down migration:
goose postgres postgres://<userName>:<password>@localhost:5432/rssagg down

# postgresql
Restarting server:
pg_ctl -D "C:\Program Files\PostgreSQL\12\data" restart