# 2016-12-20

upper_bound = 1000000
max_len = 0
st = 2
(2..upper_bound).each do |start_num|
  n = start_num
  it = 1
  while n != 1 do
    if n % 2 == 0
      n = n/2
    else
      n = 3*n+1
    end
    it += 1
  end
  if it > max_len
    max_len = it
    st = start_num
  end
end
puts st