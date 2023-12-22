CREATE TABLE IF NOT EXISTS public.marketing_leads (
    id UUID DEFAULT public.uuid_generate_v4() CONSTRAINT marketing_leads_pk PRIMARY KEY,
    product_category VArCHAR NOT NULL,
    full_name VARCHAR NOT NULL,
    company_name VARCHAR NOT NULL,
    address VARCHAR NULL,
    email VARCHAR NOT NULL,
    phone_number VARCHAR NOT NULL,
    province VARCHAR NULL,
    city VARCHAR NULL,
    district VARCHAR NULL,
    registered_date TIMESTAMP WITH TIME ZONE NOT NULL,
    source_type VARCHAR NOT NULL,
    status VARCHAR NOT NULL,
    activation_status VARCHAR NOT NULL,
    follow_up_by VARCHAR NULL,
    description TEXT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS public.marketing_leads_attachments (
    id UUID DEFAULT public.uuid_generate_v4() CONSTRAINT marketing_leads_attachments_pk PRIMARY KEY,
    file_name VARCHAR NULL,
    path VARCHAR NULL,
    mime_type VARCHAR NULL,
	marketing_leads_id uuid CONSTRAINT marketing_leads_attachment_marketing_leads_id_marketing_leads_id_fk REFERENCES public.marketing_leads NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

ALTER TABLE public.marketing_leads ADD COLUMN channel_id uuid CONSTRAINT marketing_leads_channel_id_master_channel_new_id_fk REFERENCES public.master_channel_new NULL;
ALTER TABLE public.marketing_leads ADD COLUMN template_id uuid CONSTRAINT marketing_leads_template_id_master_channel_template_id_fk REFERENCES public.master_channel_template NULL;