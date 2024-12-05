Mail-Cote : 코딩테스트 메일 전송 서비스

Apache Airflow badge Amazon S3 badge Amazos AWS badge prometheus badge grafana badge python badge Flask badge

Mail-Cote 는 grpc와 GKE(Google Kubernetes Engine)을 이용한 코딩테스트 메일 전송 서비스입니다.

매일 7시 사용자의 난이도에 맞는 문제를 큐레이팅해 이메일로 전송해주는 서비스입니다.


🔍 Preview
http://mail-cote.site/ 접속

사용자의 이메일, 비밀번호, 난이도 작성.
이후 7시 정각 난이도에 해당하는 문제 전송.

🧑🏻‍💻 Members
팀 8조입니다! grpc와 GKE를 적극적으로 활용해보고 이를 적용시키고자 노력했습니다.

허윤지 송성훈			
			
:-:	:-:		
허윤지 송성훈			
			

⌘ Project BackGround

배경 : 코딩테스트 문제를 등한시 하는 경우가 많고, 문제를 자주 푸는 습관을 기르기가 어렵다.

목표 : 코딩테스트를 자주 접함으로서, 코테 실력 향상과 습관을 기를 수 있습니다.


⚙️ Service Architecture
<img width="1049" alt="스크린샷 2024-12-05 오후 3 41 10" src="https://github.com/user-attachments/assets/540b1aa7-011e-43fd-a881-48f23c595458">
1️⃣ grpc를 통한 모듈간의 통신
grpc는 원격 모듈의 함수를 마치 자신의 모듈에서 함수를 호출하는 것처럼 하는 원격 프로시져 프로토콜이다. 
이를 통해 모듈의 DB연결을 최소화하고, 원격 모듈의 함수사용을 통해 메인 서버의 처리 속도 증가.


2️⃣ GKE를 통한 pod 관리
...dockerfile, replicaset,... 등


3️⃣ GCS bucket을 이용한 크롤링 데이터 관리
GCS: 사이트 외부 위치의 서버에 디지털 데이터가 저장되는 컴퓨터 데이터 스토리지 모드.
이를 활용해 크롤링한 데이터를 .json으로 저장 및 관리

