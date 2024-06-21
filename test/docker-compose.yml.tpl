version: "3.7"

x-logging:
  &default-logging
  options:
    max-size: "10M"
    max-file: "10"
  driver: json-file

services:
  sas_sync:
    image: localhost:5001/sas_sync:2.11.0.0
    restart: always
    container_name: sas_sync
    environment:
      - TZ=Asia/Shanghai
      - sdl_env=production
      - debug=true
    #ports:
    #  - "9002:9001"
    volumes:
      - /home/moresec/logs/sas_sync:/app/logs
      - /home/moresec/server/sas_sync/conf/:/app/conf # 配置文件位置
      - /etc/localtime:/etc/localtime
    networks:
      - sast-net
    logging: *default-logging
  sas_sca:
    image: localhost:5001/sas_sca:2.11.0.0
    restart: always
    container_name: sas_sca
    environment:
      - TZ=Asia/Shanghai
    #ports:
    #- "9007:9001"
    volumes:
      - /home/moresec/logs/sas_sca:/app/logs
      - /home/moresec/server/sas_sca/conf/:/app/conf # 配置文件位置
      - /etc/localtime:/etc/localtime
    networks:
      - sast-net
    logging: *default-logging
  sas_llm:
    image: localhost:5001/sas_llm:2.11.0.0
    restart: always
    container_name: sas_llm
    #ports:
    #   - "9004:9002"
    environment:
      - TZ=Asia/Shanghai
    volumes:
      - /home/moresec/logs/sas_llm:/app/logs
      - /home/moresec/server/sas_llm/conf/:/app/conf # 配置文件位置
      - /etc/localtime:/etc/localtime
      - /home/moresec/data/com_res/sock:/var/run
    networks:
      - sast-net
    logging: *default-logging
  scc_cve_search:
    image: localhost:5001/scc_cve_search_sast:1.2.6.0
    restart: always
    container_name: scc_cve_search
    environment:
      - sdl_env=production
      - db_host=/workspace/Top10000.db
      - debug=true
    #ports:
      #- "5005:5005"
    networks:
      - sast-net
    logging: *default-logging

  sas_engine:
    image: localhost:5001/sas_engine:2.11.0.1
    deploy:
      resources:
        limits:
          cpus: '{{ ENGINE_CPU_NUM }}'
          memory: '{{ ENGINE_MEM_NUM }}'
    restart: always
    container_name: sas_engine
    #ports:
       #- "9003:9001"
    environment:
      - TZ=Asia/Shanghai
      - sdl_env=production
    volumes:
      - /home/moresec/server/sas_engine/conf:/app/conf # 配置文件
      - /home/moresec/logs/sas_engine/:/app/logs # 日志
      - /home/moresec/server/sas_engine/data/engine-lib:/app/data/engine-lib # jar包
      - /home/moresec/server/sas_engine/data/sonar:/app/data/sonar # sonar包
      - /home/moresec/data/srv_res/sas_engine/repository:/app/data/repository #  sonar mvn
      #- /home/moresec/data/srv_res/sas_engine/scan_code:/app/scan_code #  代码
      - /home/moresec/data/mid_res/ms_code/code:/app/data/code #  代码
      - /home/moresec/data/com_res/ms_check/cpi:/app/data/ms_check/cpi
      - /home/moresec/data/srv_res/sas_engine/output:/app/data/output #  代码
      - /etc/localtime:/etc/localtime
      - /var/run/docker.sock:/var/run/docker.sock

    networks:
      - sast-net
    logging: *default-logging

  sas_sonar:
    image: localhost:5001/sas_sonar:2.11.0.0
    deploy:
      resources:
        limits:
          cpus: '2'
          memory: '4G'
    restart: always
    security_opt:
      - seccomp:unconfined # 适配低版本docker，不然会报错
    container_name: sonar.domain.com
    hostname: sonar.domain.com
    #ports:
      #- "127.0.0.1:9000:9000"
      #- "9033:9000"
    volumes:
      - /home/moresec/server/sas_sonar/conf/plugin/sonar-cxx-plugin-1.3.0.1746.jar:/opt/sonarqube/extensions/plugins/sonar-cxx-plugin-1.3.0.1746.jar
      - /home/moresec/server/sas_sonar/conf/plugin/sonar-findbugs-plugin-4.0.0.jar:/opt/sonarqube/extensions/plugins/sonar-findbugs-plugin-4.0.0.jar
      - /home/moresec/server/sas_sonar/conf/plugin/sonar-l10n-zh-plugin-1.28.jar:/opt/sonarqube/extensions/plugins/sonar-l10n-zh-plugin-1.28.jar
      - /home/moresec/logs/sas_sonar:/opt/sonarqube/logs
      - /etc/localtime:/etc/localtime
    environment:
      SONARQUBE_JDBC_URL: jdbc:postgresql://postgresql:5432/sonar
      SONARQUBE_JDBC_USERNAME: sonar
      SONARQUBE_JDBC_PASSWORD: sonar
      TZ: Asia/Shanghai
    networks:
      - sast-net
    logging: *default-logging

networks:
  sast-net:
    external:
      name: sast-net
