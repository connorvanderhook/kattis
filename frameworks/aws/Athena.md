# Athena

Athena is serverless, no infra, setup or maintenance. You pay for the queries you run (Like BQ). Athena scales automagically-executing queries in parallel-so results are fast, even on large datasets and complex queries.

You can run ad-hoc queries using ANSI SQL, w/o need to aggregate or load data into Athena.


Use Athena to generate reports, explore data w/ BI tools, or SQL clients connected with a JDBC or ODBC driver.

### Data Formats

Athena helps you analyze unstructured, semi-structured and structured data stored in S3. 

Examples:
    - CSV
    - JSON
    - Columnar formats: Apache Parquet, Apache ORC

### Integrations

- Athena **integrates** with **Amazon QuickSight** for easy data visualization.
- Athena **does not** integrate with **Amazon RDS**.
- Athena **integrates** with the **AWS Glue Data Catalog**, which ofers a **persistent metadata store** for your data in Amazon S3. 
    - Allows you to create tables and query data in Athena 
    - Central metadata store available throughout your AWS account
    - Integrated with the ETL and data discovery features of AWS Glue.
