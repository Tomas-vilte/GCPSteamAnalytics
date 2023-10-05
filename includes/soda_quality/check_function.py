def check(scan_name, checks_subpath=None, data_source="games", project_root="/opt/airflow/includes"):
    from soda.scan import Scan
    print("Running Soda Scan...")
    config_file = f"{project_root}/soda_quality/configuration.yml"
    checks_path = f"{project_root}/soda_quality/checks"

    if checks_path:
        checks_path += f"/{checks_subpath}"

    scan = Scan()
    scan.set_verbose()
    scan.add_configuration_yaml_file(config_file)
    scan.set_data_source_name(data_source)
    scan.add_sodacl_yaml_files(checks_path)
    scan.set_scan_definition_name(scan_name)
    scan.execute()
    print(scan.get_logs_text())

    if scan.has_check_fails():
        raise ValueError(f"Soda Scan failed with errors!")
    else:
        print("Soda scan successful")
        return 0

