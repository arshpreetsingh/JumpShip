data = [1,2,3,4,5,6]
main_list = []
sub_list = []

for i in range(2):
    sub_list.append(data[i])
    main_list.append(sub_list)
    sub_list = []

print(main_list)
