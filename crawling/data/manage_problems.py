import json
import random
from google.cloud import storage

# GCS 클라이언트 설정
client = storage.Client()
bucket_name = "problem-storage"  # GCS 버킷 이름
bucket = client.get_bucket(bucket_name)

# problems.json에 있는 문제 데이터를 읽어서 GCS에 업로드
def upload_problems(problems_file):

    with open(problems_file, "r", encoding="utf-8") as file:
        problems = json.load(file)

    for problem in problems:
        field = problem["field"]  # 난이도
        quiz_num = problem["quiz_num"]  # 문제 번호

        # GCS 경로 설정 (quizzes/field/quiz_num.json)
        gcs_path = f"quizzes/{field}/{quiz_num}.json"

        # JSON 데이터를 문자열로 변환
        json_data = json.dumps(problem, ensure_ascii=False, indent=4)

        # GCS에 업로드
        blob = bucket.blob(gcs_path)
        blob.upload_from_string(json_data, content_type="application/json")
        print(f"Uploaded: {gcs_path}")


# GCS에서 특정 난이도의 문제를 랜덤으로 선택함. JSON 데이터를 읽고 출력
def get_random_problem(field):
    """
    특정 난이도의 문제를 GCS에서 랜덤으로 가져옴
    - field: 문제 난이도 (예: "bronze_5")
    """
    blobs = list(bucket.list_blobs(prefix=f"quizzes/{field}/"))
    if not blobs:
        print(f"No problems found for field: {field}")
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

    # 테스트
    # 3. 특정 난이도에서 랜덤 문제 가져오기
    field = "브론즈 5"  # 난이도 선택
    print(f"\nFetching a random problem from field: {field}...")
    random_problem = get_random_problem(field)
    if random_problem:
        print(f"Random problem: {random_problem['quiz_title']}")



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