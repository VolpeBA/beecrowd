id_funcionario = IO.gets("") |> String.trim() |> String.to_integer
horas_trabalhadas = IO.gets("") |> String.trim() |> String.to_integer
valor_hora = IO.gets("") |> String.trim() |> String.to_float

salario = (valor_hora * horas_trabalhadas)

IO.puts "NUMBER = #{id_funcionario}"
IO.puts "SALARY = U$ #{:erlang.float_to_binary(salario, decimals: 2)}"
