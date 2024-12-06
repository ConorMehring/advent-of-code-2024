puzzle=$(cat input.txt)

left=()
right=()
switch="false"
for i in $puzzle; do
    if [[ $switch = "false" ]]; then
        switch="true"
        left+=($i)
    else
        switch="false"
        right+=($i)
    fi
done

# test1=(1 2 3)

# test2=(4 5 6)

# test3=(${test1[@]} ${test2[@]})

# for i in ${test3[@]}; do
#     echo $i
# done

echo $((5/2))
