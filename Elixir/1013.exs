valores = IO.gets("") |> String.trim()

vetor = String.split(valores, " ") |> Enum.map(&String.trim/1)
a = Enum.at(vetor, 0) |> String.to_integer
b = Enum.at(vetor, 1) |> String.to_integer
c = Enum.at(vetor, 2) |> String.to_integer

maiorAB = (a + b + abs(a - b)) / 2
maiorABC = ((maiorAB + c + abs(maiorAB - c)) / 2.0)

IO.puts "#{:erlang.float_to_binary(maiorABC, decimals: 0)} eh o maior"
