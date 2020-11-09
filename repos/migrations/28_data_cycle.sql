create index blocks_meta_cycle_index
	on blocks (meta_cycle desc);

CREATE FUNCTION tezos.cycles(integer, integer)
   RETURNS TABLE(cycle integer, cycle_start TIMESTAMP WITHOUT TIME ZONE, cycle_end TIMESTAMP WITHOUT TIME ZONE)
   LANGUAGE plpgsql
AS $$
    declare
    start TIMESTAMP WITHOUT TIME ZONE;
    last_block_time TIMESTAMP WITHOUT TIME ZONE;
    cycle_position integer;
    BEGIN

    select min(timestamp), max(meta_cycle_position), max(timestamp) into start, cycle_position, last_block_time from tezos.blocks where meta_cycle = $1;
    cycle := $1;
    cycle_start := start;
    cycle_end := last_block_time + (4096 - cycle_position -1) * '1 minute'::interval;
    return next;

    FOR l_counter IN $1+1..$2
    LOOP
      cycle := l_counter;
      cycle_start :=  cycle_end + '1 minute'::interval;
      cycle_end := cycle_start + interval '2 days 20 hours 16 minutes';
    return next;
    END LOOP;
END $$;

CREATE TABLE tezos.cycle_periods (
 cycle integer,
 cycle_start TIMESTAMP WITHOUT TIME ZONE,
 cycle_end TIMESTAMP WITHOUT TIME ZONE);

ALTER TABLE tezos.cycle_periods
	add constraint cycle_periods_pk
		primary key (cycle);

CREATE VIEW tezos.cycle_periods_view AS
select * from cycle_periods
UNION
select * from tezos.cycles( (select max(cycle) from cycle_periods), (select max(cycle) + 6 from cycle_periods) )
order by cycle desc;