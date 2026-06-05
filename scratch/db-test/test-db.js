const { Client } = require('pg');
const client = new Client({
  connectionString: 'postgresql://postgres:mysecretpassword@mifsaka.com:8200/kreatifdms?sslmode=disable'
});

async function main() {
  await client.connect();
  console.log("Connected to DB");
  
  const deptsRes = await client.query('SELECT id, name FROM departments');
  console.log("DEPARTMENTS:");
  console.table(deptsRes.rows);

  const racksRes = await client.query('SELECT id, name, department_id, map_pos_x, map_pos_y FROM racks');
  console.log("RACKS:");
  console.table(racksRes.rows);

  await client.end();
}

main().catch(console.error);
