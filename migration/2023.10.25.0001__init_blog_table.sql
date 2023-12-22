CREATE TABLE IF NOT EXISTS public.blog_categories (
    id UUID DEFAULT public.uuid_generate_v4() CONSTRAINT blog_categories_pk PRIMARY KEY,
    name VARCHAR NOT NULL,
    slug VARCHAR NOT NULL,
    is_active boolean NOT NULL,
    description TEXT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS public.blog_banners (
    id UUID DEFAULT public.uuid_generate_v4() CONSTRAINT blog_banners_pk PRIMARY KEY,
    name VARCHAR NOT NULL,
    file_name VARCHAR NULL,
    path VARCHAR NULL,
    url VARCHAR NULL,
    mime_type VARCHAR NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

INSERT INTO public.blog_banners (name) 
    VALUES ('BANNER_HEADER'),
        ('BANNER_FOOTER');

CREATE TABLE IF NOT EXISTS public.blog_articles (
    id UUID DEFAULT public.uuid_generate_v4() CONSTRAINT blog_articles_pk PRIMARY KEY,
    title VARCHAR NOT NULL,
	blog_category_id uuid CONSTRAINT blog_articles_blog_category_id_blog_categories_id_fk REFERENCES public.blog_categories NOT NULL,
    visibility VARCHAR NOT NULL,
    publish_date TIMESTAMP WITH TIME ZONE NOT NULL,
    content TEXT NULL,
    seo_title VARCHAR NULL,
    seo_slug VARCHAR NULL,
    seo_keywords jsonb NULL,
    seo_meta_description TEXT NULL,
    total_views numeric (15) default 0,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_blog_articles_title ON public.blog_articles (title);
CREATE INDEX idx_blog_articles_seo_title ON public.blog_articles (seo_title);
CREATE UNIQUE INDEX
  IF NOT EXISTS uidx_blog_articles_seo_slug
  ON public.blog_articles (seo_slug);

CREATE TABLE IF NOT EXISTS public.blog_articles_attachments (
    id UUID DEFAULT public.uuid_generate_v4() CONSTRAINT blog_articles_attachments_pk PRIMARY KEY,
    name VARCHAR NOT NULL,
    file_name VARCHAR NULL,
    path VARCHAR NULL,
    url VARCHAR NULL,
    mime_type VARCHAR NULL,
	blog_articles_id uuid CONSTRAINT blog_articles_attachment_blog_articles_id_blog_articles_id_fk REFERENCES public.blog_articles NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS public.blog_articles_unique_visitors (
    id UUID DEFAULT public.uuid_generate_v4() CONSTRAINT blog_blog_articles_unique_visitors_pk PRIMARY KEY,
	blog_articles_id uuid CONSTRAINT blog_blog_articles_unique_visitors_blog_articles_id_blog_articles_id_fk REFERENCES public.blog_articles NOT NULL,
    ip_address VARCHAR NOT NULL,
    last_visited_date DATE NOT NULL,
    last_visited_datetime TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX
  IF NOT EXISTS uidx_blog_articles_uv_id_ip_last_visited_date
  ON public.blog_articles_unique_visitors
  USING BTREE (blog_articles_id, ip_address, last_visited_date);