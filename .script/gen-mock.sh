#!/bin/bash

skip_files=('main.go' 'wire.go' 'wire_gen.go' 'server.go' 'router.go' 'provider.go' 'vo.go' 'entity.go')

files=$(find ./ -type f -name "*.go" ! -name "*_mock.go" ! -path "./.*/*" ! -name "*_test.go" $(printf "! -name %s " ${skip_files[@]}))

for file in ${files}; do
  # 元のファイルパスからmockディレクトリ用のパスを作成
  mock_dir="./mock$(dirname ${file} | cut -c 2-)"
  mock_file="${mock_dir}/$(basename ${file} | sed 's/\.go$/_mock.go/')"
  # mockディレクトリを作成
  mkdir -p ${mock_dir}
  # パッケージ名を取得
  p_name=$(head -n 1 ${file} | cut -c 9-)

  echo "generating ${mock_file}" 
  mockgen -source ${file} -destination ${mock_file} -package ${p_name}_mock

  # ファイルが生成されたが、行数が5行しかない場合は削除
  if [ $(grep -c '' ${mock_file}) = 5 ]; then 
    rm -rf ${mock_file}
  fi 

  echo "generated ${mock_file}" &
done

wait