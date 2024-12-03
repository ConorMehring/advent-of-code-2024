puzzle=$(cat input.txt)
match="mul([0-9]*,[0-9]*)"

split=$(echo $puzzle | sed "s/\(${match}\)/\n\1\n/g")

sum="0"
enable="true"
for item in $split; do
    if [[ $item =~ "don't()" ]]; then
        enable="false"
    elif [[ $item =~ "do()" ]]; then
        enable="true"
    fi
    if [[ $item = $match ]] && [[ $enable = "true" ]]; then
        sum="${sum}+$(echo $item | sed 's/mul(\([0-9]*,[0-9]*\))/\1/;s/,/*/')"
    fi
done

echo $(($sum))
