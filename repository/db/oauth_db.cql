// Show Cassandra cluster keyspace
DESCRIBE KEYSPACES;

// Creates a new Cassandra keyspace
// SimpleStrategy for one data center & one rack.
// Replication_factor of 1 means there's only one copy of each row in the Cassandra cluster
CREATE KEYSPACE IF NOT EXISTS oauth WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': 1};

// Select keyspace
USE oauth;