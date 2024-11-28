import json
import random
from google.cloud import storage

# GCS 클라이언트 설정
client = storage.Client()
bucket_name = "mail-cote-bucket"  # GCS 버킷 이름
bucket = client.get_bucket(bucket_name)


# problems.json에 있는 문제 데이터를 읽어서 GCS에 업로드
def upload_problems(problems_file):

    with open(problems_file, "r", encoding="utf-8") as file:
        problems = json.load(file)

    for problem in problems:
        field = problem["field"]  # 난이도
        quiz_num = problem["quiz_num"]  # 문제 번호

        # field를 공백으로 나누어 'Bronze', '5'로 분리
        grade, level = field.split()

        # GCS 경로 설정 (problems/grade/level/quiz_num.json)
        gcs_path = f"problems/{grade}/{level}/{quiz_num}.json"

        # JSON 데이터를 문자열로 변환
        json_data = json.dumps(problem, ensure_ascii=False, indent=4)

        # GCS에 업로드
        blob = bucket.blob(gcs_path)
        blob.upload_from_string(json_data, content_type="application/json")
        print(f"Uploaded: {gcs_path}")


        
# GCS에서 특정 난이도의 문제를 랜덤으로 선택하는 함수
def get_random_problem(field):

    grade, level = field.split()

    blobs = list(bucket.list_blobs(prefix=f"problems/{grade}/{level}"))
    if not blobs:
        print(f"해당 난이도에 문제가 없음: {field}")
        return None

    # 랜덤으로 문제 선택
    random_blob = random.choice(blobs)
    data = random_blob.download_as_text()
    return json.loads(data)


if __name__ == "__main__":
    
    # 1. 업로드할 문제 데이터 파일 경로
    problems_file = "problems.json"  # 로컬에 저장된 JSON 데이터 파일

    # 2. 문제 데이터를 GCS에 업로드
    print("Uploading problems to GCS...")
    upload_problems(problems_file)

    # ***** 테스트 *****
    # 3. 특정 난이도에서 랜덤 문제 가져오기
    field = "Bronze 5"  # 난이도 선택
    print(f"\n랜덤 문제를 가져올 난이도: {field}...")
    random_problem = get_random_problem(field)
    if random_problem:
        print(f"가져온 랜덤 문제: {random_problem['quiz_title']}({random_problem['quiz_num']})")




'''
# GCS 접근 권한 테스트!!!!!!!!
from google.cloud import storage

client = storage.Client()

# 버킷 목록 출력
buckets = list(client.list_buckets())
print("Buckets accessible by this service account:")
for bucket in buckets:
    print(bucket.name)

'''