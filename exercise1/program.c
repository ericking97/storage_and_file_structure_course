#include <stdio.h>
#include <stdlib.h>

#define FILENAME "dump.txt"

struct Record
{
  int id;
  char name[20];
};

int main()
{
  struct Record records[] = {
      {1, "Juan"},
      {2, "Paco"},
      {3, "Pedro"},
      {4, "Ana"},
      {5, "Rocío"}};

  // Write records on file
  FILE *fpw; // Stands for file pointer write
  if ((fpw = fopen(FILENAME, "w")) == NULL)
  {
    printf("Error al abrir o crear el archivo.\nSe forzará el cierre");
    return EXIT_FAILURE;
  }
  fwrite(records, sizeof(struct Record), 5, fpw);
  fclose(fpw);

  // Read records from file
  FILE *fpr; // Stands for file pointer read
  if ((fpr = fopen(FILENAME, "r")) == NULL)
  {
    printf("Error al abrir el archivo.\nSe forzará el cierre");
    return EXIT_FAILURE;
  }
  while (!feof(fpr))
  {
    struct Record record;
    fread(&record, sizeof(struct Record), 1, fpr);
    printf("ID: %d\nNombre: %s\n", record.id, record.name);
  }
  fclose(fpr);

  return EXIT_SUCCESS;
}
