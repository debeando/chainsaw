
mysql -h 127.0.0.1 -u root -padmin demo -Bse "SELECT id FROM foo ORDER BY RAND() LIMIT 70000" | sort -n > ids.txt
