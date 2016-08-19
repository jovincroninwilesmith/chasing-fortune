CREATE TABLE btc_market_data (
    volume decimal,
   	latest_trade integer,
   	bid integer,
	high decimal,
	currency varchar(60),
	currency_volume decimal,
	ask decimal,
	close decimal,
	avg decimal,
	symbol varchar(60),
	low decimal,
	tstamp timestamptz DEFAULT now()	
);

