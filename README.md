# MySQL Dump Table Purger

A command line application to remove table data from MySQL dumps, while keeping the table structure itself intact.

Useful if your MySQL dumps / backups include large tables you don't want to spend time restoring to your development environment - for example, logs or other non-vital data.

## Usage / Commands

### List Tables in Dump

    Lists all tables found in the given MySQL dump.

    ```bash
    $ mysqldumptablepurger list <inputFile>
    ```

### Remove Tables

    Removes the content of the given table names from the dump, and writes the resulting file to the `outputFile` path given.

    ```bash
    $ mysqldumptablepurger remove <inputFile> <outputFile> <table1> <table2> <...>
    ```

### Remove Tables - Using Config File

    Removes the content of tables using a predefined YAML config file, useful if you want to maintain a list of tables to frequently remove.

    ```bash
    $ mysqldumptablepurger remove-by-conf <inputFile> <configFile>
    ```

## Config File Example

    ```yaml
    outputFile: mysql_dump_purged.sql
    tables:
        - user_logs
        - admin_logs
    ```