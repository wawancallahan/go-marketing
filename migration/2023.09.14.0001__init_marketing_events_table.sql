CREATE TABLE IF NOT EXISTS public.marketing_events (
    id UUID DEFAULT public.uuid_generate_v4() CONSTRAINT marketing_events_pk PRIMARY KEY,
    event_name VARCHAR NOT NULL,
    event_time TIMESTAMP WITH TIME ZONE NOT NULL,
    event_location VARCHAR NULL,
    event_type VARCHAR NOT NULL,
    channel_event VARCHAR NOT NULL,
    measurement_event VARCHAR NOT NULL,
    status VARCHAR NOT NULL,
    province VARCHAR NULL,
    city VARCHAR NULL,
    participant NUMERIC(12, 0) NOT NULL DEFAULT 0,
    pic_name VARCHAR NULL,
    support_name VARCHAR NULL,
    created_by_id uuid CONSTRAINT marketing_leads_created_by_id_users_id_fk REFERENCES pm_auth_internal.users NULL,
    updated_by_id uuid CONSTRAINT marketing_leads_updated_by_id_users_id_fk REFERENCES pm_auth_internal.users NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);