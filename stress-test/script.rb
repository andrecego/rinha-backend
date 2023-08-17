# {"apelido": "hWgACCtvJYiNbeCwQmXNbiGPseNsx", "nome": "qXuQbPwgtoNuJGAJqLQOEyRwrCejKAUfSUOKhNJcvDDveZGKMBPekQQEdkbJnxEjQcNJAldQdJXr", "nascimento": "1941-03-18", "stack": ["Swift", "Ruby", "Go", "Java", "Rust", "Swift", "Postgres", "Javascript", "Perl", "Java", "C"]}
File.readlines('stress-test/data/pessoas-payloads.tsv').each do |line|
  `curl -X POST -H "Content-Type: application/json" -d '#{line}' http://localhost:3333/pessoas`
end
