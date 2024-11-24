# JSON 데이터를 저장, 로드하는 유틸리티 함수가 포함된 파일
import json

def save_to_json(data, file_name):
    with open(file_name, 'w', encoding='utf-8') as f:
        json.dump(data, f, ensure_ascii=False, indent=4)

def load_json(file_name):
    try:
        with open(file_name, 'r', encoding='utf-8') as f:
            return json.load(f)
    except FileNotFoundError:
        return []
