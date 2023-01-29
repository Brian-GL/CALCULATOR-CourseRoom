%Función que nos ayuda a obtener un nuevo individuo recalculado si el
%individuo pasado como parámetro sobrepasa algunos de los límites establecidos en alguna dimensión de este.

%En donde requerimos como parámetro los siguientes valores:

%Individuo: El individuo a penalizar.
%Mínimo: El límite mínimo permitido.
%Máximo: El límite máximo permitido.
%Dimensión: El número de dimensiones que el individuo tiene.

%Retornamos un nuevo individuo recalculado si, y solo si, el individuo sobrepasa, en alguna dimensión, un límite.


function IndividuoRecalculado = Penalizacion(Individuo,Minimo,Maximo, Dimension)

    %Creamos e inicializamos el valor del nuevo individuo recalculado:
    IndividuoRecalculado = zeros(Dimension,1); %Tomará el mismo tamaño que el individuo.
    
    %Creamos un for para iterar cada dimensión del Individuo.
    for i = 1: Dimension
        if Minimo(i) > Individuo(i)  || Individuo(i) > Maximo(i) %Si el individuo sobrepasa algún limite, entonces en IndividuoRecalculado recalculamos una nueva posición válida:
            IndividuoRecalculado(i) = randi([Minimo(i) Maximo(i)]);
        else %Sino, pasamos el valor del individuo tal cual pero casteado a entero (ya que como recordaremos, tenemos solamente posiciones enteras en las imágenes):
            IndividuoRecalculado(i) = int32(Individuo(i));
        end
    end

end