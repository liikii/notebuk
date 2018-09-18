#include <stdio.h>
#include <stdlib.h>


int main(int argc, char *argv[])
{
    int i;

    double* p;    // We uses this reference variable to access
                  // dynamically created array elements

    p = calloc(10, sizeof(double) );  // Make double array of 10 elements

    for ( i = 0; i < 10; i++ )
       *(p + i) = i;            // put value i in array element i

    for ( i = 0; i < 10; i++ )
       printf("*(p + %d) = %lf\n", i, *(p+i) );


    free(p);      // Un-reserve the first array

    putchar('\n');

    p = calloc(4, sizeof(double) );  // Make a NEW double array of 4 elements

    for ( i = 0; i < 4; i++ )
       *(p + i) = i*i;            // put value i*i in array element i

    for ( i = 0; i < 4; i++ )
       printf("*(p + %d) = %lf\n", i, *(p+i) );

    free(p);      // Un-reserve the second array
}
