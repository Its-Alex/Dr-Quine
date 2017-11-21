#include <stdio.h>
/*
	commentaire1
*/
char *getA() { return "#include <stdio.h>%c/*%c%ccommentaire1%c*/%cchar *getA() { return %c%s%c; }%cint main(){%c/*%c%ccommentaire2%c*/%cprintf(getA(), 10, 10, 9, 10, 10, 34, getA(), 34, 10, 10, 10, 9, 10, 10, 10);%c}"; }
int main(){
/*
	commentaire2
*/
printf(getA(), 10, 10, 9, 10, 10, 34, getA(), 34, 10, 10, 10, 9, 10, 10, 10);
}