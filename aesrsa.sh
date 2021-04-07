#!/bin/bash
function getTiming(){
        start=$1
        end=$2

        start_s=`echo $start| cut -d '.' -f 1`
        start_ns=`echo $start| cut -d '.' -f 2`
        end_s=`echo $end| cut -d '.' -f 1`
        end_ns=`echo $end| cut -d '.' -f 2`

        time_micro=$(( (10#$end_s-10#$start_s)*1000000 + (10#$end_ns/1000 - 10#$start_ns/1000) ))
        time_ms=`expr $time_micro/1000  | bc `

        echo "$time_micro microseconds"
        echo "$time_ms ms"
}

j=$1

# 生成RSA私钥
# openssl genrsa -out rsa.key 2048
# 从私钥中提取公钥
# openssl rsa -in rsa.key -pubout -out rsa_pub.key

begin_time=`date +%s.%N`
for ((i=1; i<=j; i++))
do
# 公钥加密
openssl rsautl -encrypt -in plain.txt -inkey rsa_pub.key -pubin -out cipher.txt
# 私钥解密
openssl rsautl -decrypt -in cipher.txt -inkey rsa.key -out plain2.txt
done
end_time=`date +%s.%N`
getTiming  $begin_time $end_time

begin_time=`date +%s.%N`
# 对称加密
for ((i=1; i<=j; i++))
do
openssl enc -e -aes-128-cbc -a -iter 100 -pbkdf2 -in plain.txt -out cipher2.txt -k 123
openssl enc -d -aes-128-cbc -a -iter 100 -pbkdf2 -in cipher2.txt -out plain3.txt.txt -k 123
done
end_time=`date +%s.%N`
getTiming  $begin_time $end_time

