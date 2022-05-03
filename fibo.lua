function fibo(x --[[int]], y --[[int]], z --[[int]])
  fibonacci = {y, z} --int
  print(string.format("1: %d\n2: %d", fibonacci[1], fibonacci[2]))
  for i=3 --[[int]], x do
    table.insert(fibonacci, fibonacci[i-1] + fibonacci[i-2])
    print(string.format("%d: %d", i, fibonacci[i]))
  end
  phi=fibonacci[#fibonacci]/fibonacci[#fibonacci-1]
  print(string.format("\n\nPhi: %f", phi))
end

function main()
  if #arg ~= 3 then
    print("You have to provide 3 arguments.")
    return 1
  else
    print(fibo(tonumber(arg[3]) --[[int]], tonumber(arg[1]) --[[int]], tonumber(arg[2]) --[[int]]))
    return 0
  end
end
main()
