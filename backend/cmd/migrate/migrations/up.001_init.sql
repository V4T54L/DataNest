CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255),
    password VARCHAR(255) NOT NULL
);

CREATE TABLE dashboards (
    dashboard_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    dashboard_name VARCHAR(255) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE charts (
    chart_id SERIAL PRIMARY KEY,
    dashboard_id INT NOT NULL,
    chart_type VARCHAR(50) NOT NULL, -- Type of the chart (e.g., bar, pie, line, etc.)
    chart_data JSONB NOT NULL, -- The actual chart data stored as JSON (e.g., the dataset)
    FOREIGN KEY (dashboard_id) REFERENCES dashboards(dashboard_id)
);
