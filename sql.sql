create TABLE cards (
    id BIGINT generated by DEFAULT as IDENTITY,
    cards TEXT not null,
    type text not null,
    agent bigint
)

CREATE TABLE agent (
    id BIGINT generated by DEFAULT as IDENTITY,
    agent text not null
)