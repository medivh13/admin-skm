CREATE TABLE master.jawabans (
    id bigint NOT NULL,
    pertanyaan_id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    pilihan_satu character varying(255) DEFAULT '' NOT NULL,
    pilihan_dua character varying(255) DEFAULT '' NOT NULL,
    pilihan_tiga character varying(255) DEFAULT '' NOT NULL,
    pilihan_empat character varying(255) DEFAULT '' NOT NULL,
    CONSTRAINT fk_pertanyaans
      FOREIGN KEY(pertanyaan_id) 
	  REFERENCES master.pertanyaans(id)
);

CREATE SEQUENCE master.jawabans_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE master.jawabans_id_seq OWNED BY master.jawabans.id;

ALTER TABLE ONLY master.jawabans ALTER COLUMN id SET DEFAULT nextval('master.jawabans_id_seq'::regclass);

ALTER TABLE ONLY master.jawabans
    ADD CONSTRAINT jawabans_pkey PRIMARY KEY (id);

CREATE INDEX jawabans_index ON master.jawabans USING btree (deleted_at, id, pertanyaan_id);


insert into master.jawabans (pertanyaan_id, pilihan_satu, pilihan_dua, pilihan_tiga, pilihan_empat)
Values (1, 'jawabansatu', 'jawabandua', 'jawabantiga', 'jawabanempat');
insert into master.jawabans (pertanyaan_id, pilihan_satu, pilihan_dua, pilihan_tiga, pilihan_empat)
Values (2, 'jawabansatu', 'jawabandua', 'jawabantiga', 'jawabanempat');
insert into master.jawabans (pertanyaan_id, pilihan_satu, pilihan_dua, pilihan_tiga, pilihan_empat)
Values (3, 'jawabansatu', 'jawabandua', 'jawabantiga', 'jawabanempat');
insert into master.jawabans (pertanyaan_id, pilihan_satu, pilihan_dua, pilihan_tiga, pilihan_empat)
Values (4, 'jawabansatu', 'jawabandua', 'jawabantiga', 'jawabanempat');
insert into master.jawabans (pertanyaan_id, pilihan_satu, pilihan_dua, pilihan_tiga, pilihan_empat)
Values (5, 'jawabansatu', 'jawabandua', 'jawabantiga', 'jawabanempat');
insert into master.jawabans (pertanyaan_id, pilihan_satu, pilihan_dua, pilihan_tiga, pilihan_empat)
Values (6, 'jawabansatu', 'jawabandua', 'jawabantiga', 'jawabanempat');
insert into master.jawabans (pertanyaan_id, pilihan_satu, pilihan_dua, pilihan_tiga, pilihan_empat)
Values (7, 'jawabansatu', 'jawabandua', 'jawabantiga', 'jawabanempat');
insert into master.jawabans (pertanyaan_id, pilihan_satu, pilihan_dua, pilihan_tiga, pilihan_empat)
Values (8, 'jawabansatu', 'jawabandua', 'jawabantiga', 'jawabanempat');