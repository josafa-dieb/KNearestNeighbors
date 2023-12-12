# KNearestNeighbors
# Equipe
- Jaziel Loureiro
- Naum Josafá
- Samuel Jonas

### Dataset

 - [Dataset Iris.csv](https://github.com/josafa-dieb/KNearestNeighbors/blob/main/dataset/Iris.csv)
 - [Resultado: dados de treinamento](https://github.com/josafa-dieb/KNearestNeighbors/tree/main/result)
 - [Resultado: dados de testes](https://github.com/josafa-dieb/KNearestNeighbors/tree/main/result)
 

### Dataset Iris
- a base de dados dos valores das flores é composto por 3 (Íris-setosa, Íris-virgínica e Íris-versicolor) espécies diferentes com seus respectivos valores.
- o resultados de testes e treinamentos foram selecionados aleatóriamente, e foi utilizada a distância euclidiana dos elemento em teste com os elementos do treinamento, e assim foram feitas as classificações.

![Contagem de Species](https://github.com/josafa-dieb/KNearestNeighbors/assets/49986895/0056ea49-9d76-4208-8983-c5b9a20d56fe)

|Id |SepalLengthCm|SepalWidthCm|PetalLengthCm|PetalWidthCm|Species        |
|---|-------------|------------|-------------|------------|---------------|
|1  |5.1          |3.5         |1.4          |0.2         |Iris-setosa    |
|2  |4.9          |3.0         |1.4          |0.2         |Iris-setosa    |
|3  |4.7          |3.2         |1.3          |0.2         |Iris-setosa    |
|4  |4.6          |3.1         |1.5          |0.2         |Iris-setosa    |
|5  |5.0          |3.6         |1.4          |0.2         |Iris-setosa    |
|6  |5.4          |3.9         |1.7          |0.4         |Iris-setosa    |
|7  |4.6          |3.4         |1.4          |0.3         |Iris-setosa    |
|8  |5.0          |3.4         |1.5          |0.2         |Iris-setosa    |
|9  |4.4          |2.9         |1.4          |0.2         |Iris-setosa    |
|10 |4.9          |3.1         |1.5          |0.1         |Iris-setosa    |
|11 |5.4          |3.7         |1.5          |0.2         |Iris-setosa    |
|12 |4.8          |3.4         |1.6          |0.2         |Iris-setosa    |
|13 |4.8          |3.0         |1.4          |0.1         |Iris-setosa    |
|14 |4.3          |3.0         |1.1          |0.1         |Iris-setosa    |
|15 |5.8          |4.0         |1.2          |0.2         |Iris-setosa    |
|16 |5.7          |4.4         |1.5          |0.4         |Iris-setosa    |
|17 |5.4          |3.9         |1.3          |0.4         |Iris-setosa    |
|18 |5.1          |3.5         |1.4          |0.3         |Iris-setosa    |
|19 |5.7          |3.8         |1.7          |0.3         |Iris-setosa    |
|20 |5.1          |3.8         |1.5          |0.3         |Iris-setosa    |
|21 |5.4          |3.4         |1.7          |0.2         |Iris-setosa    |
|22 |5.1          |3.7         |1.5          |0.4         |Iris-setosa    |
|23 |4.6          |3.6         |1.0          |0.2         |Iris-setosa    |
|24 |5.1          |3.3         |1.7          |0.5         |Iris-setosa    |
|25 |4.8          |3.4         |1.9          |0.2         |Iris-setosa    |
|26 |5.0          |3.0         |1.6          |0.2         |Iris-setosa    |
|27 |5.0          |3.4         |1.6          |0.4         |Iris-setosa    |
|28 |5.2          |3.5         |1.5          |0.2         |Iris-setosa    |
|29 |5.2          |3.4         |1.4          |0.2         |Iris-setosa    |
|30 |4.7          |3.2         |1.6          |0.2         |Iris-setosa    |
|31 |4.8          |3.1         |1.6          |0.2         |Iris-setosa    |
|32 |5.4          |3.4         |1.5          |0.4         |Iris-setosa    |
|33 |5.2          |4.1         |1.5          |0.1         |Iris-setosa    |
|34 |5.5          |4.2         |1.4          |0.2         |Iris-setosa    |
|35 |4.9          |3.1         |1.5          |0.1         |Iris-setosa    |
|36 |5.0          |3.2         |1.2          |0.2         |Iris-setosa    |
|37 |5.5          |3.5         |1.3          |0.2         |Iris-setosa    |
|38 |4.9          |3.1         |1.5          |0.1         |Iris-setosa    |
|39 |4.4          |3.0         |1.3          |0.2         |Iris-setosa    |
|40 |5.1          |3.4         |1.5          |0.2         |Iris-setosa    |
|41 |5.0          |3.5         |1.3          |0.3         |Iris-setosa    |
|42 |4.5          |2.3         |1.3          |0.3         |Iris-setosa    |
|43 |4.4          |3.2         |1.3          |0.2         |Iris-setosa    |
|44 |5.0          |3.5         |1.6          |0.6         |Iris-setosa    |
|45 |5.1          |3.8         |1.9          |0.4         |Iris-setosa    |
|46 |4.8          |3.0         |1.4          |0.3         |Iris-setosa    |
|47 |5.1          |3.8         |1.6          |0.2         |Iris-setosa    |
|48 |4.6          |3.2         |1.4          |0.2         |Iris-setosa    |
|49 |5.3          |3.7         |1.5          |0.2         |Iris-setosa    |
|50 |5.0          |3.3         |1.4          |0.2         |Iris-setosa    |
|51 |7.0          |3.2         |4.7          |1.4         |Iris-versicolor|
|52 |6.4          |3.2         |4.5          |1.5         |Iris-versicolor|
|53 |6.9          |3.1         |4.9          |1.5         |Iris-versicolor|
|54 |5.5          |2.3         |4.0          |1.3         |Iris-versicolor|
|55 |6.5          |2.8         |4.6          |1.5         |Iris-versicolor|
|56 |5.7          |2.8         |4.5          |1.3         |Iris-versicolor|
|57 |6.3          |3.3         |4.7          |1.6         |Iris-versicolor|
|58 |4.9          |2.4         |3.3          |1.0         |Iris-versicolor|
|59 |6.6          |2.9         |4.6          |1.3         |Iris-versicolor|
|60 |5.2          |2.7         |3.9          |1.4         |Iris-versicolor|
|61 |5.0          |2.0         |3.5          |1.0         |Iris-versicolor|
|62 |5.9          |3.0         |4.2          |1.5         |Iris-versicolor|
|63 |6.0          |2.2         |4.0          |1.0         |Iris-versicolor|
|64 |6.1          |2.9         |4.7          |1.4         |Iris-versicolor|
|65 |5.6          |2.9         |3.6          |1.3         |Iris-versicolor|
|66 |6.7          |3.1         |4.4          |1.4         |Iris-versicolor|
|67 |5.6          |3.0         |4.5          |1.5         |Iris-versicolor|
|68 |5.8          |2.7         |4.1          |1.0         |Iris-versicolor|
|69 |6.2          |2.2         |4.5          |1.5         |Iris-versicolor|
|70 |5.6          |2.5         |3.9          |1.1         |Iris-versicolor|
|71 |5.9          |3.2         |4.8          |1.8         |Iris-versicolor|
|72 |6.1          |2.8         |4.0          |1.3         |Iris-versicolor|
|73 |6.3          |2.5         |4.9          |1.5         |Iris-versicolor|
|74 |6.1          |2.8         |4.7          |1.2         |Iris-versicolor|
|75 |6.4          |2.9         |4.3          |1.3         |Iris-versicolor|
|76 |6.6          |3.0         |4.4          |1.4         |Iris-versicolor|
|77 |6.8          |2.8         |4.8          |1.4         |Iris-versicolor|
|78 |6.7          |3.0         |5.0          |1.7         |Iris-versicolor|
|79 |6.0          |2.9         |4.5          |1.5         |Iris-versicolor|
|80 |5.7          |2.6         |3.5          |1.0         |Iris-versicolor|
|81 |5.5          |2.4         |3.8          |1.1         |Iris-versicolor|
|82 |5.5          |2.4         |3.7          |1.0         |Iris-versicolor|
|83 |5.8          |2.7         |3.9          |1.2         |Iris-versicolor|
|84 |6.0          |2.7         |5.1          |1.6         |Iris-versicolor|
|85 |5.4          |3.0         |4.5          |1.5         |Iris-versicolor|
|86 |6.0          |3.4         |4.5          |1.6         |Iris-versicolor|
|87 |6.7          |3.1         |4.7          |1.5         |Iris-versicolor|
|88 |6.3          |2.3         |4.4          |1.3         |Iris-versicolor|
|89 |5.6          |3.0         |4.1          |1.3         |Iris-versicolor|
|90 |5.5          |2.5         |4.0          |1.3         |Iris-versicolor|
|91 |5.5          |2.6         |4.4          |1.2         |Iris-versicolor|
|92 |6.1          |3.0         |4.6          |1.4         |Iris-versicolor|
|93 |5.8          |2.6         |4.0          |1.2         |Iris-versicolor|
|94 |5.0          |2.3         |3.3          |1.0         |Iris-versicolor|
|95 |5.6          |2.7         |4.2          |1.3         |Iris-versicolor|
|96 |5.7          |3.0         |4.2          |1.2         |Iris-versicolor|
|97 |5.7          |2.9         |4.2          |1.3         |Iris-versicolor|
|98 |6.2          |2.9         |4.3          |1.3         |Iris-versicolor|
|99 |5.1          |2.5         |3.0          |1.1         |Iris-versicolor|
|100|5.7          |2.8         |4.1          |1.3         |Iris-versicolor|
|101|6.3          |3.3         |6.0          |2.5         |Iris-virginica |
|102|5.8          |2.7         |5.1          |1.9         |Iris-virginica |
|103|7.1          |3.0         |5.9          |2.1         |Iris-virginica |
|104|6.3          |2.9         |5.6          |1.8         |Iris-virginica |
|105|6.5          |3.0         |5.8          |2.2         |Iris-virginica |
|106|7.6          |3.0         |6.6          |2.1         |Iris-virginica |
|107|4.9          |2.5         |4.5          |1.7         |Iris-virginica |
|108|7.3          |2.9         |6.3          |1.8         |Iris-virginica |
|109|6.7          |2.5         |5.8          |1.8         |Iris-virginica |
|110|7.2          |3.6         |6.1          |2.5         |Iris-virginica |
|111|6.5          |3.2         |5.1          |2.0         |Iris-virginica |
|112|6.4          |2.7         |5.3          |1.9         |Iris-virginica |
|113|6.8          |3.0         |5.5          |2.1         |Iris-virginica |
|114|5.7          |2.5         |5.0          |2.0         |Iris-virginica |
|115|5.8          |2.8         |5.1          |2.4         |Iris-virginica |
|116|6.4          |3.2         |5.3          |2.3         |Iris-virginica |
|117|6.5          |3.0         |5.5          |1.8         |Iris-virginica |
|118|7.7          |3.8         |6.7          |2.2         |Iris-virginica |
|119|7.7          |2.6         |6.9          |2.3         |Iris-virginica |
|120|6.0          |2.2         |5.0          |1.5         |Iris-virginica |
|121|6.9          |3.2         |5.7          |2.3         |Iris-virginica |
|122|5.6          |2.8         |4.9          |2.0         |Iris-virginica |
|123|7.7          |2.8         |6.7          |2.0         |Iris-virginica |
|124|6.3          |2.7         |4.9          |1.8         |Iris-virginica |
|125|6.7          |3.3         |5.7          |2.1         |Iris-virginica |
|126|7.2          |3.2         |6.0          |1.8         |Iris-virginica |
|127|6.2          |2.8         |4.8          |1.8         |Iris-virginica |
|128|6.1          |3.0         |4.9          |1.8         |Iris-virginica |
|129|6.4          |2.8         |5.6          |2.1         |Iris-virginica |
|130|7.2          |3.0         |5.8          |1.6         |Iris-virginica |
|131|7.4          |2.8         |6.1          |1.9         |Iris-virginica |
|132|7.9          |3.8         |6.4          |2.0         |Iris-virginica |
|133|6.4          |2.8         |5.6          |2.2         |Iris-virginica |
|134|6.3          |2.8         |5.1          |1.5         |Iris-virginica |
|135|6.1          |2.6         |5.6          |1.4         |Iris-virginica |
|136|7.7          |3.0         |6.1          |2.3         |Iris-virginica |
|137|6.3          |3.4         |5.6          |2.4         |Iris-virginica |
|138|6.4          |3.1         |5.5          |1.8         |Iris-virginica |
|139|6.0          |3.0         |4.8          |1.8         |Iris-virginica |
|140|6.9          |3.1         |5.4          |2.1         |Iris-virginica |
|141|6.7          |3.1         |5.6          |2.4         |Iris-virginica |
|142|6.9          |3.1         |5.1          |2.3         |Iris-virginica |
|143|5.8          |2.7         |5.1          |1.9         |Iris-virginica |
|144|6.8          |3.2         |5.9          |2.3         |Iris-virginica |
|145|6.7          |3.3         |5.7          |2.5         |Iris-virginica |
|146|6.7          |3.0         |5.2          |2.3         |Iris-virginica |
|147|6.3          |2.5         |5.0          |1.9         |Iris-virginica |
|148|6.5          |3.0         |5.2          |2.0         |Iris-virginica |
|149|6.2          |3.4         |5.4          |2.3         |Iris-virginica |
|150|5.9          |3.0         |5.1          |1.8         |Iris-virginica |

### Seleção Treinamento (80%)
- seleção de dados para treinamento foi de 80% em relação ao dataset original
![Contagem de Species](https://github.com/josafa-dieb/KNearestNeighbors/assets/49986895/68aace99-7f68-4e49-9c1f-697ab66dc101)

|Id |SepalLengthCm|SepalWidthCm|PetalLengthCm|PetalWidthCm|Species        |
|---|-------------|------------|-------------|------------|---------------|
|28 |5.20         |3.50        |1.50         |0.20        |Iris-setosa    |
|21 |5.40         |3.40        |1.70         |0.20        |Iris-setosa    |
|49 |5.30         |3.70        |1.50         |0.20        |Iris-setosa    |
|74 |6.10         |2.80        |4.70         |1.20        |Iris-versicolor|
|58 |4.90         |2.40        |3.30         |1.00        |Iris-versicolor|
|53 |6.90         |3.10        |4.90         |1.50        |Iris-versicolor|
|43 |4.40         |3.20        |1.30         |0.20        |Iris-setosa    |
|33 |5.20         |4.10        |1.50         |0.10        |Iris-setosa    |
|121|6.90         |3.20        |5.70         |2.30        |Iris-virginica |
|68 |5.80         |2.70        |4.10         |1.00        |Iris-versicolor|
|1  |5.10         |3.50        |1.40         |0.20        |Iris-setosa    |
|81 |5.50         |2.40        |3.80         |1.10        |Iris-versicolor|
|135|6.10         |2.60        |5.60         |1.40        |Iris-virginica |
|110|7.20         |3.60        |6.10         |2.50        |Iris-virginica |
|146|6.70         |3.00        |5.20         |2.30        |Iris-virginica |
|61 |5.00         |2.00        |3.50         |1.00        |Iris-versicolor|
|147|6.30         |2.50        |5.00         |1.90        |Iris-virginica |
|99 |5.10         |2.50        |3.00         |1.10        |Iris-versicolor|
|20 |5.10         |3.80        |1.50         |0.30        |Iris-setosa    |
|45 |5.10         |3.80        |1.90         |0.40        |Iris-setosa    |
|84 |6.00         |2.70        |5.10         |1.60        |Iris-versicolor|
|136|7.70         |3.00        |6.10         |2.30        |Iris-virginica |
|50 |5.00         |3.30        |1.40         |0.20        |Iris-setosa    |
|18 |5.10         |3.50        |1.40         |0.30        |Iris-setosa    |
|102|5.80         |2.70        |5.10         |1.90        |Iris-virginica |
|82 |5.50         |2.40        |3.70         |1.00        |Iris-versicolor|
|87 |6.70         |3.10        |4.70         |1.50        |Iris-versicolor|
|126|7.20         |3.20        |6.00         |1.80        |Iris-virginica |
|107|4.90         |2.50        |4.50         |1.70        |Iris-virginica |
|15 |5.80         |4.00        |1.20         |0.20        |Iris-setosa    |
|65 |5.60         |2.90        |3.60         |1.30        |Iris-versicolor|
|54 |5.50         |2.30        |4.00         |1.30        |Iris-versicolor|
|77 |6.80         |2.80        |4.80         |1.40        |Iris-versicolor|
|130|7.20         |3.00        |5.80         |1.60        |Iris-virginica |
|23 |4.60         |3.60        |1.00         |0.20        |Iris-setosa    |
|140|6.90         |3.10        |5.40         |2.10        |Iris-virginica |
|137|6.30         |3.40        |5.60         |2.40        |Iris-virginica |
|90 |5.50         |2.50        |4.00         |1.30        |Iris-versicolor|
|19 |5.70         |3.80        |1.70         |0.30        |Iris-setosa    |
|59 |6.60         |2.90        |4.60         |1.30        |Iris-versicolor|
|8  |5.00         |3.40        |1.50         |0.20        |Iris-setosa    |
|52 |6.40         |3.20        |4.50         |1.50        |Iris-versicolor|
|67 |5.60         |3.00        |4.50         |1.50        |Iris-versicolor|
|25 |4.80         |3.40        |1.90         |0.20        |Iris-setosa    |
|113|6.80         |3.00        |5.50         |2.10        |Iris-virginica |
|34 |5.50         |4.20        |1.40         |0.20        |Iris-setosa    |
|30 |4.70         |3.20        |1.60         |0.20        |Iris-setosa    |
|112|6.40         |2.70        |5.30         |1.90        |Iris-virginica |
|2  |4.90         |3.00        |1.40         |0.20        |Iris-setosa    |
|141|6.70         |3.10        |5.60         |2.40        |Iris-virginica |
|5  |5.00         |3.60        |1.40         |0.20        |Iris-setosa    |
|11 |5.40         |3.70        |1.50         |0.20        |Iris-setosa    |
|41 |5.00         |3.50        |1.30         |0.30        |Iris-setosa    |
|75 |6.40         |2.90        |4.30         |1.30        |Iris-versicolor|
|93 |5.80         |2.60        |4.00         |1.20        |Iris-versicolor|
|64 |6.10         |2.90        |4.70         |1.40        |Iris-versicolor|
|144|6.80         |3.20        |5.90         |2.30        |Iris-virginica |
|106|7.60         |3.00        |6.60         |2.10        |Iris-virginica |
|116|6.40         |3.20        |5.30         |2.30        |Iris-virginica |
|92 |6.10         |3.00        |4.60         |1.40        |Iris-versicolor|
|123|7.70         |2.80        |6.70         |2.00        |Iris-virginica |
|97 |5.70         |2.90        |4.20         |1.30        |Iris-versicolor|
|117|6.50         |3.00        |5.50         |1.80        |Iris-virginica |
|66 |6.70         |3.10        |4.40         |1.40        |Iris-versicolor|
|63 |6.00         |2.20        |4.00         |1.00        |Iris-versicolor|
|131|7.40         |2.80        |6.10         |1.90        |Iris-virginica |
|120|6.00         |2.20        |5.00         |1.50        |Iris-virginica |
|72 |6.10         |2.80        |4.00         |1.30        |Iris-versicolor|
|38 |4.90         |3.10        |1.50         |0.10        |Iris-setosa    |
|39 |4.40         |3.00        |1.30         |0.20        |Iris-setosa    |
|48 |4.60         |3.20        |1.40         |0.20        |Iris-setosa    |
|37 |5.50         |3.50        |1.30         |0.20        |Iris-setosa    |
|27 |5.00         |3.40        |1.60         |0.40        |Iris-setosa    |
|86 |6.00         |3.40        |4.50         |1.60        |Iris-versicolor|
|22 |5.10         |3.70        |1.50         |0.40        |Iris-setosa    |
|124|6.30         |2.70        |4.90         |1.80        |Iris-virginica |
|40 |5.10         |3.40        |1.50         |0.20        |Iris-setosa    |
|145|6.70         |3.30        |5.70         |2.50        |Iris-virginica |
|122|5.60         |2.80        |4.90         |2.00        |Iris-virginica |
|118|7.70         |3.80        |6.70         |2.20        |Iris-virginica |
|148|6.50         |3.00        |5.20         |2.00        |Iris-virginica |
|108|7.30         |2.90        |6.30         |1.80        |Iris-virginica |
|128|6.10         |3.00        |4.90         |1.80        |Iris-virginica |
|143|5.80         |2.70        |5.10         |1.90        |Iris-virginica |
|31 |4.80         |3.10        |1.60         |0.20        |Iris-setosa    |
|115|5.80         |2.80        |5.10         |2.40        |Iris-virginica |
|3  |4.70         |3.20        |1.30         |0.20        |Iris-setosa    |
|91 |5.50         |2.60        |4.40         |1.20        |Iris-versicolor|
|101|6.30         |3.30        |6.00         |2.50        |Iris-virginica |
|26 |5.00         |3.00        |1.60         |0.20        |Iris-setosa    |
|149|6.20         |3.40        |5.40         |2.30        |Iris-virginica |
|127|6.20         |2.80        |4.80         |1.80        |Iris-virginica |
|73 |6.30         |2.50        |4.90         |1.50        |Iris-versicolor|
|119|7.70         |2.60        |6.90         |2.30        |Iris-virginica |
|35 |4.90         |3.10        |1.50         |0.10        |Iris-setosa    |
|24 |5.10         |3.30        |1.70         |0.50        |Iris-setosa    |
|79 |6.00         |2.90        |4.50         |1.50        |Iris-versicolor|
|56 |5.70         |2.80        |4.50         |1.30        |Iris-versicolor|
|4  |4.60         |3.10        |1.50         |0.20        |Iris-setosa    |
|12 |4.80         |3.40        |1.60         |0.20        |Iris-setosa    |
|85 |5.40         |3.00        |4.50         |1.50        |Iris-versicolor|
|44 |5.00         |3.50        |1.60         |0.60        |Iris-setosa    |
|71 |5.90         |3.20        |4.80         |1.80        |Iris-versicolor|
|29 |5.20         |3.40        |1.40         |0.20        |Iris-setosa    |
|125|6.70         |3.30        |5.70         |2.10        |Iris-virginica |
|62 |5.90         |3.00        |4.20         |1.50        |Iris-versicolor|
|57 |6.30         |3.30        |4.70         |1.60        |Iris-versicolor|
|6  |5.40         |3.90        |1.70         |0.40        |Iris-setosa    |
|109|6.70         |2.50        |5.80         |1.80        |Iris-virginica |
|142|6.90         |3.10        |5.10         |2.30        |Iris-virginica |
|47 |5.10         |3.80        |1.60         |0.20        |Iris-setosa    |
|150|5.90         |3.00        |5.10         |1.80        |Iris-virginica |
|83 |5.80         |2.70        |3.90         |1.20        |Iris-versicolor|
|14 |4.30         |3.00        |1.10         |0.10        |Iris-setosa    |
|105|6.50         |3.00        |5.80         |2.20        |Iris-virginica |
|60 |5.20         |2.70        |3.90         |1.40        |Iris-versicolor|
|129|6.40         |2.80        |5.60         |2.10        |Iris-virginica |
|94 |5.00         |2.30        |3.30         |1.00        |Iris-versicolor|
|100|5.70         |2.80        |4.10         |1.30        |Iris-versicolor|
|42 |4.50         |2.30        |1.30         |0.30        |Iris-setosa    |

### Seleção Testes (20%)
- o percentual de seleção de dados para testes foi de 20% em relação ao dataset original
![Contagem Testes](https://github.com/josafa-dieb/KNearestNeighbors/assets/49986895/df088006-3b4e-4bad-942e-8e1607959481)

|Id |SepalLengthCm|SepalWidthCm|PetalLengthCm|PetalWidthCm|Species        |
|---|-------------|------------|-------------|------------|---------------|
|0  |0.00         |0.00        |0.00         |0.00        |Species        |
|7  |4.60         |3.40        |1.40         |0.30        |Iris-setosa    |
|9  |4.40         |2.90        |1.40         |0.20        |Iris-setosa    |
|10 |4.90         |3.10        |1.50         |0.10        |Iris-setosa    |
|13 |4.80         |3.00        |1.40         |0.10        |Iris-setosa    |
|16 |5.70         |4.40        |1.50         |0.40        |Iris-setosa    |
|17 |5.40         |3.90        |1.30         |0.40        |Iris-setosa    |
|32 |5.40         |3.40        |1.50         |0.40        |Iris-setosa    |
|36 |5.00         |3.20        |1.20         |0.20        |Iris-setosa    |
|46 |4.80         |3.00        |1.40         |0.30        |Iris-setosa    |
|51 |7.00         |3.20        |4.70         |1.40        |Iris-versicolor|
|55 |6.50         |2.80        |4.60         |1.50        |Iris-versicolor|
|69 |6.20         |2.20        |4.50         |1.50        |Iris-versicolor|
|70 |5.60         |2.50        |3.90         |1.10        |Iris-versicolor|
|76 |6.60         |3.00        |4.40         |1.40        |Iris-versicolor|
|78 |6.70         |3.00        |5.00         |1.70        |Iris-versicolor|
|80 |5.70         |2.60        |3.50         |1.00        |Iris-versicolor|
|88 |6.30         |2.30        |4.40         |1.30        |Iris-versicolor|
|89 |5.60         |3.00        |4.10         |1.30        |Iris-versicolor|
|95 |5.60         |2.70        |4.20         |1.30        |Iris-versicolor|
|96 |5.70         |3.00        |4.20         |1.20        |Iris-versicolor|
|98 |6.20         |2.90        |4.30         |1.30        |Iris-versicolor|
|103|7.10         |3.00        |5.90         |2.10        |Iris-virginica |
|104|6.30         |2.90        |5.60         |1.80        |Iris-virginica |
|111|6.50         |3.20        |5.10         |2.00        |Iris-virginica |
|114|5.70         |2.50        |5.00         |2.00        |Iris-virginica |
|132|7.90         |3.80        |6.40         |2.00        |Iris-virginica |
|133|6.40         |2.80        |5.60         |2.20        |Iris-virginica |
|134|6.30         |2.80        |5.10         |1.50        |Iris-virginica |
|138|6.40         |3.10        |5.50         |1.80        |Iris-virginica |
|139|6.00         |3.00        |4.80         |1.80        |Iris-virginica |

## Resultados

|      Métricas | Resultados |
| ----------------- | ---------------------------------------------------------------- |
| Acurácia | 0.94 |
| Precisão | 0.97 |
| Recall   | 1.00 |


## Referências
 - [kNN - uma introdução aos algoritmos de aprendizado de máquina](https://medium.com/dftblog/knn-introdu%C3%A7%C3%A3o-aos-algoritmos-de-aprendizado-de-m%C3%A1quina-dd2107693651)
