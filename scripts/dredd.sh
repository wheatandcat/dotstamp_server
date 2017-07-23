mysql -uroot stamp_test -N -e 'show tables' | xargs -IARG mysql -uroot -e 'truncate table ARG' stamp_test
mysql -u root stamp_test < ./scripts/dump.sql
