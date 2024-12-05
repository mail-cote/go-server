# 📫 **Mail-Cote**: 코딩테스트 메일 전송 서비스 📫
![KakaoTalk_Photo_2024-12-05-16-07-31](https://github.com/user-attachments/assets/91f4ac29-c0b8-49c9-81eb-998e2a1bd83a)

![GKE](https://img.shields.io/badge/GKE-blue)
![gRPC](https://img.shields.io/badge/gRPC-green)
![GCP](https://img.shields.io/badge/GCP-orange)
![Golang](https://img.shields.io/badge/Golang-blue)
![Cloud SQL](https://img.shields.io/badge/Cloud%20SQL-lightblue)
![Docker](https://img.shields.io/badge/Docker-2496ED)
![Nginx](https://img.shields.io/badge/Nginx-darkgreen)
![Streamlit](https://img.shields.io/badge/Streamlit-red)
![GCS](https://img.shields.io/badge/GCS-blue)


---

Mail-Cote는 **gRPC**와 **Google Kubernetes Engine**를 활용한 **코딩테스트 메일 전송 서비스**입니다.  
매일 **오전 7시**, 사용자의 난이도에 맞는 문제를 큐레이팅해 이메일로 전송합니다.

🔗 [**서비스 바로가기**](http://mail-cote.site/)

---

## 📋 **서비스 기능**
1. 사용자는 이메일, 비밀번호, 난이도를 입력.
2. 매일 **오전 7시**, 사용자의 난이도에 맞는 코딩테스트 문제를 이메일로 전송.
3. 큐레이션된 문제를 통해 코딩테스트 실력을 꾸준히 향상.

---

## 👩🏻‍💻 **Members**
| Name   | Role       |
|--------|------------|
| 허윤지 | Back-End   |
| 송성훈 | Back-End   |

**팀 8조**는 `gRPC`와 `GKE`를 적극적으로 활용하며 모듈 간 통신을 효율화하고 확장 가능한 서비스를 제공합니다.

---

## 🧑🏻‍💻 **프로젝트 배경**
- **배경**: 코딩테스트 문제를 등한시하거나 꾸준히 풀지 못하는 경우가 많음.
- **목표**: 코딩테스트 문제를 매일 접하며 실력 향상과 학습 습관을 형성.

---

## ⚙️ **서비스 아키텍처**
<img width="1062" alt="스크린샷 2024-12-05 오후 4 04 53" src="https://github.com/user-attachments/assets/e3e69110-c49c-46ad-a692-851f39bb3369">

### 1️⃣ **gRPC를 통한 모듈 간 통신**
- `gRPC`: 원격 모듈의 함수를 로컬처럼 호출할 수 있는 원격 프로시저 프로토콜.
- 모듈 간 DB 연결 최소화 및 처리 속도 최적화.

### 2️⃣ **GKE를 통한 Pod 관리**
- `GKE`: Kubernetes 기반의 컨테이너 오케스트레이션 플랫폼.
- Dockerfile, ReplicaSet 등을 활용하여 서비스 안정성 확보.

### 3️⃣ **GCS Bucket을 이용한 데이터 관리**
- `Google Cloud Storage`: 크롤링 데이터를 `.json` 파일로 저장 및 관리.
- 외부 서버에서 디지털 데이터를 효율적으로 관리.

### 4️⃣ **Cloud sqld을 통한 RDBMS 관리**
- 자동 백업, 및 손쉬운 복구
- Google Cloud 관리형 서비스이므로, 호율적 운영 및 확장 가능성 측면에서 좋음.
---

## 🛠️ **Tech Stack**
- **Language**: Python, Go
- **Framework**: gRPC
- **Infrastructure**: GKE, Docker, GCS, Cloud sql
- **Plus**: Nginx, Streamlit

---

## 🌟 **프로젝트 미리보기**
서비스 미리보기 페이지: [**Mail-Cote**](http://mail-cote.site/)

---


## 📈 **Future Plans**
1. 코딩테스트 문제 추천 알고리즘 고도화.
2. UI/UX 개선 및 대시보드 추가.
3. 다양한 코딩테스트 플랫폼 연동(GitHub Copilot, LeetCode, etc).
4. 기능 확장
