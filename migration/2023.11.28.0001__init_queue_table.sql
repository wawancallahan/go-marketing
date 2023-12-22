CREATE TABLE IF NOT EXISTS public.queue_logs (
    id UUID DEFAULT public.uuid_generate_v4() CONSTRAINT pj_queue_logs_pk PRIMARY KEY,
    exchange VARCHAR NULL,
    routing_key VARCHAR NULL,
    payload JSONB NULL,
    headers JSONB NULL,
    flow VARCHAR NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);