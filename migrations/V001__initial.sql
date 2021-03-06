BEGIN;
CREATE TABLE IF NOT EXISTS "invitees"
(
    "id"                     bigint GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    "name"                   varchar                                 NOT NULL,
    "is_child"               boolean                                 NULL     DEFAULT false,
    "has_plus_one"           boolean                                 NOT NULL DEFAULT false,
    "is_bridesmaid"          boolean                                 NOT NULL DEFAULT false,
    "is_groomsman"           boolean                                 NOT NULL DEFAULT false,
    "plus_one_name"          varchar                                 NULL,
    "phone"                  varchar                                 NULL,
    "email"                  varchar                                 NULL,
    "address_line_1"         varchar                                 NULL,
    "address_line_2"         varchar                                 NULL,
    "address_city"           varchar                                 NULL,
    "address_state"          varchar                                 NULL,
    "address_postal_code"    varchar                                 NULL,
    "address_country"        varchar                                 NULL,
    "rsvp_response"          boolean                                 NULL,
    "invitee_party_invitees" bigint                                  NULL,
    PRIMARY KEY ("id")
);
CREATE TABLE IF NOT EXISTS "invitee_parties"
(
    "id"   bigint GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    "name" varchar                                 NOT NULL,
    "code" varchar UNIQUE                          NOT NULL,
    PRIMARY KEY ("id")
);
ALTER TABLE "invitees"
    ADD CONSTRAINT "invitees_invitee_parties_invitees" FOREIGN KEY ("invitee_party_invitees") REFERENCES "invitee_parties" ("id") ON DELETE SET NULL;
COMMIT;
