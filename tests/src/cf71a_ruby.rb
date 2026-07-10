def main
  input = $stdin.read.split
  return if input.empty?
  n = input[0].to_i
  (1..n).each do |i|
    s = input[i]
    if s.length > 10
      puts "#{s[0]}#{s.length - 2}#{s[-1]}"
    else
      puts s
    end
  end
end

main
