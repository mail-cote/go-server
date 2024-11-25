# 문제 리스트와 상세 정보를 수집하는 함수가 포함된 파일
import logging
import requests
from bs4 import BeautifulSoup, NavigableString

logging.basicConfig(level=logging.INFO)

def get_problem_list(page_url):
    logging.info(f"문제 리스트 from {page_url}")
    headers = {
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
        "Referer": "https://www.acmicpc.net/"
    }
    response = requests.get(page_url, headers=headers)

    soup = BeautifulSoup(response.text, 'html.parser')

    problems = []
    
    # 모든 <tr> 태그 가져오기
    rows = soup.select("#problemset tbody tr")

    for row in rows:
        try:
            # 문제 번호 (td[1])
            quiz_num = row.select_one("td:nth-of-type(1)").text.strip()
            # 문제 제목 및 URL (td[2]/a)
            quiz_title_element = row.select_one("td:nth-of-type(2) a")
            quiz_title = quiz_title_element.text.strip()
            quiz_url = f"https://www.acmicpc.net/problem/{quiz_num}"

            problems.append({
                "quiz_num": quiz_num,
                "quiz_title": quiz_title,
                "url": quiz_url
            })

        except AttributeError as e:
            print(f"예외 발생: {e}")
            break
    
    print(f"수집된 문제 수: {len(problems)}")

    return problems


def get_problem_details(problem_url):

    headers = {
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
        "Referer": "https://www.acmicpc.net/"
    }

    response = requests.get(problem_url, headers=headers)  
    response.raise_for_status()
    soup = BeautifulSoup(response.text, 'html.parser')

    # selector로 HTML 요소를 찾아 텍스트와 이미지 src를 순서대로 조합.
    def get_text_and_images(selector):
        element = soup.select_one(selector)
        if not element:
            return None

        # 텍스트와 이미지를 순서대로 포함
        
        result = []
        for child in element.descendants:
            if isinstance(child, NavigableString):
                # 텍스트 노드 처리
                text = child.strip()
                if text:
                    result.append(text)
            elif child.name == 'img':
                # 이미지 노드 처리
                img_src = child.get('src')
                if img_src:
                    result.append(f"[Image: {img_src}]")

        return "\n".join(result)


    return {
        "description": get_text_and_images('#problem_description'),
        "input_desc": get_text_and_images('#problem_input'),
        "output_desc": get_text_and_images('#problem_output'),
        "input_ex": get_text_and_images('#sample-input-1'),
        "output_ex": get_text_and_images('#sample-output-1'),
        "time_limit": soup.select_one('#problem-info tbody tr td:nth-of-type(1)').text.strip(),
        "memory_limit": soup.select_one('#problem-info tbody tr td:nth-of-type(2)').text.strip(),
    }