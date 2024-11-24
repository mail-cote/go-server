# 모든 크롤링 로직 연결
# utils, config 폴더에서 필요한 모듈을 불러와 실행함

from utils.bs_utils import get_page_urls
from utils.scraping_utils import get_problem_list, get_problem_details
from utils.json_utils import save_to_json
from utils.logging_utils import setup_logger
from config.settings import START_URLS
import os

setup_logger()

def main():
    all_data = []
    for difficulty in START_URLS:
        field = difficulty['field']
        page_urls = get_page_urls(difficulty['url'])
        
        for page_url in page_urls:
            problems = get_problem_list(page_url)

            for problem in problems:
                details = get_problem_details(problem['url'])
                problem_data = {**problem, **details, "field": field}
                all_data.append(problem_data)
                save_to_json(all_data, os.path.join("data", "problems.json"))
    
    print("***************크롤링 끝 *************")
    print(f"수집된 문제 개수: {len(all_data)}")

if __name__ == "__main__":
    main()
