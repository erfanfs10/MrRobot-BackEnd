package queries

import "fmt"

const (
	Base = `
		SELECT
			p.id,
			p.title,
			p.slug,
			p.content,
			p.excerpt,
			p.image,
			p.published_at,
			p.meta_title,
			p.meta_description,
			json_agg(DISTINCT jsonb_build_object('title', c.title)) AS categories,
			json_agg(DISTINCT jsonb_build_object('title', t.title)) AS tags
		FROM
			posts            AS p
		LEFT JOIN
			post_categories  AS pc ON p.id = pc.post_id
		LEFT JOIN
			categories       AS c  ON c.id = pc.category_id
		LEFT JOIN
			post_tags        AS pt ON p.id = pt.post_id
		LEFT JOIN
			tags             AS t  ON t.id = pt.tag_id
		WHERE
			p.status = 'published' %s
		GROUP BY
			p.id, p.title, p.slug, p.published_at, p.excerpt
		%s
		%s
	`
)

func BuildPostListQuery(filterBy, title string) string {

	validWheres := map[string]string{
		"random": `%s`,
		"tag": `AND EXISTS (
					SELECT 1
					FROM post_tags ptf
					JOIN tags tf ON tf.id = ptf.tag_id
					WHERE ptf.post_id = p.id
					AND tf.title = '%s')`,

		"category": `AND EXISTS (
						SELECT 1
						FROM post_categories pcf
						JOIN categories cf ON cf.id = pcf.category_id
						WHERE pcf.post_id = p.id
						AND cf.title = '%s')`,
		"detail": `AND p.slug = '%s'`,
		"product": `AND EXISTS (
						SELECT 1
						FROM post_categories pcf
						JOIN categories cf ON cf.id = pcf.category_id
						WHERE pcf.post_id = p.id
						AND cf.title = '%s')`,
	}

	validOrderBy := map[string]string{
		"random":   `ORDER BY RANDOM()`,
		"tag":      `ORDER BY p.published_at DESC`,
		"category": `ORDER BY p.published_at DESC`,
		"detail":   `ORDER BY p.published_at DESC`,
		"product":  `ORDER BY p.published_at DESC`,
	}

	validLimit := map[string]string{
		"random":   `limit 4`,
		"tag":      "",
		"category": "",
		"detail":   "",
		"product":  "LIMIT 4",
	}

	where := validWheres[filterBy]
	orderBy := validOrderBy[filterBy]
	limit := validLimit[filterBy]

	where = fmt.Sprintf(where, title)

	return fmt.Sprintf(Base, where, orderBy, limit)
}
