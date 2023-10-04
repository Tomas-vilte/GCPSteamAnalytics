from pipeline.src.extract_data import extract_data_games_details, save_data_to_csv

if __name__ == "__main__":
    data, columns = extract_data_games_details()
    save_data_to_csv(data, columns)
