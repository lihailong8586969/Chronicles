// 打印出登录对话框
User ui_print_login_dialog() {

    char name[30];
    char password[30];

    while (1) {

        // 输入用户名
        printf("Please input your name : ");
        scanf("%s", &name);

        // 吃掉回车
        getchar();

        // 安全显示输入的密码
        printf("Please input your password : ");
        char c;
        for (int i = 0;;) {

            c = getch();

            // 如果是删除键
            if ((int) (c) == 8) {

                if(i>0){

                    i = i - 1;
                    printf("\b \b");
                }
            }
                // 如果是回车键
            else if (c == '\n' || c == '\r') {

                password[i] = '\0';
                putchar('\n');
                break;
            }
                // 不是删除键也不是回车键，则输出星号
            else {

                password[i] = c;
                putch('*');
                ++i;
            }
        }

        // 判断用户名和密码是否合法
        int len1 = strlen(name);
        int len2 = strlen(password);
        if (len1 < 5 || len1 > 20 || len2 < 5 || len2 > 20) {

            printf("Your name or password is not valid\n");
        } else {
            
            break;
        }
    }

    User user = {name, password};

    return user;
}

// 安全的输入密码
void ui_password_input_safely(char *password) {

    char c;
    for (int i = 0;;) {

        c = getch();

        // 如果是删除键
        if ((int) (c) == 8) {

            if (i > 0) {

                i = i - 1;
                printf("\b \b");
            }
        }
            // 如果是回车键
        else if (c == '\n' || c == '\r') {

            password[i] = '\0';
            putchar('\n');
            break;
        }
            // 不是删除键也不是回车键，则输出星号
        else {

            password[i] = c;
            putch('*');
            ++i;
        }
    }
}
