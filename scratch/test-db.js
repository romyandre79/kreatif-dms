const { Client } = require('pg');
const client = new Client({
  connectionString: 'postgresql://postgres:mysecretpassword@mifsaka.com:8200/kreatifdms?sslmode=disable'
});

async function main() {
  await client.connect();
  console.log("Connected to DB");
  
  const racksRes = await client.query('SELECT id, name, max_boxes_capacity, current_boxes_count FROM racks');
  console.log("RACKS:");
  console.table(racksRes.rows);

  const boxesRes = await client.query('SELECT id, rack_id, name, max_docs_capacity, current_docs_count FROM boxes');
  console.log("BOXES:");
  console.table(boxesRes.rows);

  const docsRes = await client.query('SELECT id, box_id, title, status FROM documents');
  console.log("DOCUMENTS:");
  console.table(docsRes.rows);

  await client.end();
}

main().catch(console.error);
