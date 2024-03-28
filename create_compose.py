services = {}
base_port = 4000  # 시작 포트 번호
port_increment = 100  # 각 서비스별 포트 증가량

for i in range(1, 6):  # testnodes1부터 testnodes10까지
    service_name = f"testnodes{i}"
    ports = [f"{base_port+j}:{4000}" for j in range(port_increment)]
    services[service_name] = {
        'image': 'testnodes',
        'container_name': service_name,
        'ports': ports,
        'environment': [
            'ServerPort=' + str(base_port)
        ],
        'restart': 'unless-stopped'
    }
    base_port += port_increment  # 다음 서비스의 시작 포트 업데이트

# Docker Compose 파일 생성
compose_file = {
    'version': '3.8',
    'services': services
}

import yaml
with open('docker-compose.yml', 'w') as file:
    yaml.dump(compose_file, file, default_flow_style=False)
