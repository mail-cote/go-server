import requests
from bs4 import BeautifulSoup

def get_page_urls(base_url):
    headers = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"
    }

    page_urls = []
    try:
        response = requests.get(base_url, headers=headers)
        response.raise_for_status()  # HTTP 요청 오류 확인
        soup = BeautifulSoup(response.text, 'html.parser')

        # 페이지네이션의 마지막 페이지 번호 가져오기
        pagination = soup.find('ul', class_='pagination')
        if pagination:
            last_page = pagination.find_all('li')[-1].find('a')
            last_page_number = int(last_page.text.strip()) if last_page else 1
        else:
            last_page_number = 1

        # 모든 페이지 URL 생성
        page_urls = [f"{base_url}&page={i}" for i in range(1, last_page_number + 1)]
    except requests.exceptions.RequestException as e:
        print(f"HTTP 요청 에러: {e}")
    except Exception as e:
        print(f"Unexpected error: {e}")

    return page_urls