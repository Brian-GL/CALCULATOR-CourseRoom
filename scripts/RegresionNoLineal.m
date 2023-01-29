function Resultado = RegresionNoLineal(registros)

    %Generamos e inicializamos primeramente los valores globales que se van a utilizar:

    %Definimos por defecto el número de dimensiones a utilizar:
    Dimensiones = 8;
    NumeroDeDatos = size(registros,1);

    %Dimensiones:

    %{

        a: Calificación de la tarea.
        b: Puntualidad de la tarea.
        c: Promedio general.
        d: Puntualidad general.
        e: Promedio general del curso.
        f: Puntualidad general del curso.
        g: varianza del promedio general.
        h: varianza de la puntualidad general.
    %}

    % Establecemos la función base:
    FuncionBase = @(a,b,c,d,e,f,g,h) ((a+b+c+d+e+f+g+h)/8) + ((a + e)*exp(2) + )
    
   

    %  Establecemos el modelo no lineal con valor g, el cual es parte fundamental de la fórmula que usaremos para obtener la regresión no lineal.
    g = @(w1,w2) w1*exp(w2*xp); 
    
    % Generamos nuestra función objetivo, obtenida y basada en la misma fórmula para obtener la regresión no lineal.
    Funcion = @(w1,w2) (1/(2*NumeroDeDatos))*sum((yp - g(w1,w2)).^2);


    
    return 1;

end


