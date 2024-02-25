CREATE TABLE public.wallet (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    balance INT DEFAULT 100
);

CREATE TABLE public.operation (
    time TIMESTAMP WITH TIME ZONE NOT NULL,
    fromId UUID NOT NULL,
    toId UUID NOT NULL,
    amount INT NOT NULL,
    CONSTRAINT fk_fromId FOREIGN KEY (fromId) REFERENCES public.wallet(id),
    CONSTRAINT fk_toId FOREIGN KEY (toId) REFERENCES public.wallet(id)
);


INSERT INTO wallet DEFAULT VALUES;
INSERT INTO wallet DEFAULT VALUES;


SELECT * FROM wallet;

INSERT INTO wallet DEFAULT VALUES RETURNING id, balance