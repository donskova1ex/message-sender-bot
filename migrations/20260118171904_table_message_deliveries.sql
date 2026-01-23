-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS message_deliveries (
    id BIGSERIAL PRIMARY KEY,
    message_id BIGINT NOT NULL REFERENCES messages(id) ON DELETE CASCADE,
    chat_id BIGINT NOT NULL,
    is_sent BOOLEAN NOT NULL DEFAULT FALSE,
    sent_at TIMESTAMP,
    error TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE INDEX IF NOT EXISTS idx_message_deliveries_unsent ON message_deliveries (message_id, chat_id) WHERE is_sent = false;
CREATE INDEX IF NOT EXISTS idx_message_deliveries_chat_id ON message_deliveries (chat_id) WHERE deleted_at IS NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS message_deliveries;
-- +goose StatementEnd
