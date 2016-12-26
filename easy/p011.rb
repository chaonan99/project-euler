# 2016-12-20

array = []
File.open("../input/p011_input.txt", "r") do |f|
  f.each_line do |line|
     array << line.split
  end
end
array = array.map{|row| row.map{|a| a.to_i}}
max_number = 0
array.each_with_index do |row, i|
  row.each_with_index do |ele, j|
    begin
      multi_row = array[i][j]*array[i+1][j]*array[i+2][j]*array[i+3][j]
      multi_col = array[i][j]*array[i][j+1]*array[i][j+2]*array[i][j+3]
      multi_dia = array[i][j]*array[i+1][j+1]*array[i+2][j+2]*array[i+3][j+3]
      multi_res = array[i+3][j]*array[i+2][j+1]*array[i+1][j+2]*array[i][j+3]
      max_number = [max_number, multi_row, multi_col, multi_dia, multi_res].max
    rescue TypeError, NoMethodError => exception_variable
    end
  end
end
puts max_number