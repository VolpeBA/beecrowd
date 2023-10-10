id_funcionario = IO.gets("") |> String.trim()
salario_fixo = IO.gets("") |> String.trim() |> String.to_float
vendas = IO.gets("") |> String.trim() |> String.to_float

salario = (salario_fixo + (vendas * 0.15))

IO.puts "TOTAL = R$ #{:erlang.float_to_binary(salario, decimals: 2)}"
