%Función que nos ayuda a realizar el algoritmo 'PolinizacionEvolutiva', el cual nos ayuda a obtener las posibles soluciones de la funci�n objetivo.

%En donde requerimos valores para los siguientes par�metros:

%Parámetros generales:

%Imagen -> La imagen original.
%Plantilla -> La imagen a buscar match en la imagen original.
%NumeroDeIteraciones -> El número de iteraciones que el algoritmo realizar� para encontrar la posible soluci�n.
%NumeroDeIndividuos -> El número de individuos a crear y a utilizar en el algoritmo.
%Minimo -> El vector de límite m�nimo de cada valor de la dimensi�n. 
%Maximo -> El vector de límite m�ximo de cada valor de la dimensi�n.
%Dimension -> El número de dimensiones que la funci�n presenta.

%Par�metros particulares:

%ParametroDePaso -> Representa el valor del par�metro de paso que los individuos utilizar�n (Valor lamdba en el algoritmo de la polinizaci�n de la flor).
%CriterioDeProbabilidad = El criterio de probabilidad que nos ayudar� a 'switchear' la polinizaci�n global y local.
%FactorDeAmplificacion -> es el factor amplificaci�n, el cual se encuentra entre el rango de valores 0 <= F <= 2 (Factor De Amplificaci�n en el algoritmo de evoluci�n diferencial).
%ConstanteDeRecombinacion -> es una constante de recombinaci�n, la cual se encuentra entre los valores [0 1]. Adem�s, esta constante nos sirve para la hora de la cruza (como en el algoritmo de evoluci�n diferencial).

%Esta funci�n retorna a Solucion, que es un vector con todos los valores �ptimos de cada dimensi�n, por lo que soluci�n es un vector de Dimensi�n x 1.
%Adem�s, retornamos a Poblaci�n ,que son todos aquellos individuos creados para el algoritmo, utilizado principalmente para graficar los resultados m�s adelante.

%Este algoritmo se basa principalmente en la combinaci�n de los m�todos que utilizan los siguientes algoritmos:
% - Polinizaci�n de la flor.
% - Evoluci�n Diferencial.


function Solucion = PolinizacionEvolutiva(Imagen,Plantilla,NumeroDeIteraciones,NumeroDeIndividuos, Minimo, Maximo, Dimension, ParametroDePaso, CriterioDeProbabilidad ,FactorDeAmplificacion,ConstanteDeRecombinacion)

    %El primer paso que tendremos que realizar es la inicializaci�n y creaci�n de todos los valores, vectores, matrices , datos, etc. que utilizaremos para optimizar la funci�n objetivo.

    %     F a s e    D e     I n i c i a l i z a c i � n     D e     V a l o r e s  %
    
    %Creamos la matriz de individuos del algoritmo, lo llamaremos Poblaci�n:
    
    %La matriz ser� de tama�o Dimension x NumeroDeIndividuos:
    Poblacion = zeros(Dimension, NumeroDeIndividuos); %Inicializamos todos los valores en cero, para desp�es llenarlos con valores aleatorios dentro del rango especificado.

    %Creamos ahora un vector para guardar solamente las evaluaciones de cada individuo de la poblaci�n en la funci�n objetivo, la llamaremos Evaluaciones:
    
    %En este caso tendremos un vector de tama�o número de individuos, ya que la evaluaci�n es un  valor escalar.
    Evaluaciones = zeros(1,NumeroDeIndividuos); %Al igual que los valores anteriores, este vector se crea con valores en cero.
    
    %Establecemos el valor de sigma2, el cual es requerido para el vuelo de Levy:
    Sigma2 = ((gamma(1+ParametroDePaso)/(ParametroDePaso*gamma((1+ParametroDePaso)/2)))*((sin(pi*ParametroDePaso/2))/(2^((ParametroDePaso-1)/2)))) ^ (1/ParametroDePaso); % Aparte, para la creaci�n de este valor usamos la funnci�n gamma que nos brinda por defecto matlab.

    %Ahora, damos paso a la inicializaci�n de los valores de la matriz de  Poblaci�n y el vector de Evaluaciones:

    %Creamos un for desde 1 hasta el número de individuos para almacenar la informaci�n respectiva de cada uno de los valores correspondientes.
    for i=1:NumeroDeIndividuos 
       
        %Creamos la variable respectiva al individuo actual:
        Individuo = zeros(Dimension,1);
        
        %Iteramos para cada dimensi�n con el objetivo de obtener la posici�n aletoria del individuo:
        for j = 1: Dimension
             %Creamos, para esa dimensi�n del individuo, un valor aleatorio
             %dentro del rango especificado en los valores Minimo y Maximo dados como par�metros.
             %Sin embargo, debemos de especificar que los valores aleatorios sean enteros (Ya que en la imagen nos posicionamos en �ndices enteros).
             Individuo(j) = randi([Minimo(j) Maximo(j)]);
         end
        
        %A�adimos el nuevo individuo a la poblaci�n 
        Poblacion(:,i) = Individuo;
        
        %Ahora, en el vector de evaluaciones guardamos la evaluaci�n de este individuo en la funci�n objetivo:
        Evaluaciones(i) = NCC(Imagen,Plantilla,Individuo);
    end
    
    %  F i n   F a s e    D e     I n i c i a l i z a c i � n     D e     V a l o r e s  %
    
    %------------------------------------------------------------------------------------%
    %       F a s e    D e    D e s a r r o l l o    D e l    A l g o r i t m o          %
    
    %Creamos un for desde 1 hasta el número de iteraciones para desarrollar el algoritmo tal cantidad de veces:
    for iteracion = 1:NumeroDeIteraciones

        
        %Generamos otro for, el que nos ayudar� a iterar a cada individuo de la poblaci�n:
        for i = 1: NumeroDeIndividuos
            
            %Establecemos primeramente un valor aleatorio r, tal que 0<=r<=1, con el prop�sito de seleccionar el tipo de polinizaci�n: global o local.
            r = rand; %Y como lo hemos ido haciendo en cada pr�ctica, lo obtenemos con rand.
            
            %Si el valor aleatorio es menor al criterio de probabilidad previamente generado, entonces, realizamos la polinizaci�n global.
            if r < CriterioDeProbabilidad
                %    *Fase De Polinizaci�n Global*
                
                %Como primer paso debemos de obtener el �ndice del mejor individuo de la poblaci�n, es decir, el que tenga la mejor evaluaci�n en la funci�n:
                IndiceDelMejor = 1;
                Comparacion = 0;
                for a=1:NumeroDeIndividuos
                   EvaluacionActual = Evaluaciones(a);
                   if EvaluacionActual > Comparacion
                      IndiceDelMejor = a;
                      Comparacion = EvaluacionActual;
                   end
                end

                %Inicializamos este tipo de polinizaci�n creando los vectores de u y v de distribuci�n normal requeridos en esta parte:    
                u = normrnd(0,Sigma2,[Dimension 1]); %Nos apoyamos utilizando la funci�n normrnd que nos brinda matlab, adem�s, utilizamos el valor previamente creado de sigma2
                v = normrnd(0,1,[Dimension 1]);%En este caso, en v aplicamos una distribuci�n normal con 1 y no con sigma2.

                %Desp�es, podemos crear ahora si nuestro vector L mediante la f�rmula utilizada en el algoritmo de Mantegna:

                L = u./ (abs(v).^(1/ParametroDePaso)); %Y como recordaremos, L es un vector y no un escalar, por lo que utilizamos los operadores '.'

                %Ya obtenidos estos valores, podemos ahora crear nuestro nuevo vector de soluci�n candidata:

                SolucionCandidata = Poblacion(:,i) + L.* (Poblacion(:,IndiceDelMejor) - Poblacion(:,i));
                
                %Mediante la funci�n de penalizaci�n, recalcularemos la soluci�n candidata obtenida con anterioridad:
                SolucionCandidata = Penalizacion(SolucionCandidata,Minimo,Maximo, Dimension);
                
                %Generamos en una variable la evaluaci�n de la nueva soluci�n candidata:
                EvaluacionSolucionCandidata = NCC(Imagen,Plantilla,SolucionCandidata);

                %Si esta nueva evaluaci�n es mejor que la evaluaci�n previa del individuo, entonces actualizamos el individuo por la soluci�n candidata, incluyendo su nueva evaluaci�n.
                if EvaluacionSolucionCandidata > Evaluaciones(i) %Recordando que es un problema de maximizaci�n en este caso.
                    Poblacion(:,i) = SolucionCandidata;
                    Evaluaciones(i) = EvaluacionSolucionCandidata;
                end

              %    *Fin Fase De Polinizaci�n Global*
              
            else %Si no, entonces realizamos evolucion diferencial, el cual se intercambi� por el m�todo de la polinizaci�n local.
                
                %Generamos otro for, para visitar a cada individuo dentro de la poblacion:
                for j = 1 : NumeroDeIndividuos
                    %Ahora, debemos de obtener valores aleatorios r1, r2 & r3, tales que no sean igual a 'j'.
                    r1 = randi([1 NumeroDeIndividuos]);
                    %Usamos while para cada r(n), con el hecho de validar que r1 != r2 != r3 != j.
                    while r1 == j
                        r1 = randi([1 NumeroDeIndividuos]);
                    end
                    r2 = randi([1 NumeroDeIndividuos]);
                    while r2 == j || r1 == r2
                        r2 = randi([1 NumeroDeIndividuos]);
                    end
                    r3 = randi([1 NumeroDeIndividuos]);
                    while r3 == j || r3 == r1 || r3 == r2
                        r3 = randi([1 NumeroDeIndividuos]);
                    end
                    %Generamos ahora el vector mutado que el algoritmo de evoluci�n diferencial utiliza [ vi := xr1 + F (xr2 - xr3) ]
                    Vector = Poblacion(:,r1) + FactorDeAmplificacion * (Poblacion(:,r2) - Poblacion(:,r3));
                    %Generamos un vector auxiliar que nos ayudar� a mutar, todo eso dependiento de nuestra constante de recombinaci�n.
                    VectorMutado = zeros(Dimension,1);
                    for k = 1:Dimension
                        %Si nuestro número random es menor o igual a nuestra constante de recombinaci�n, es as� que podemos tomar la parte de la dimensi�n 'k' del vector para combinar.
                        if rand <= ConstanteDeRecombinacion
                            VectorMutado(k) = Vector(k); 
                        else %sino, pasamos el valor del individuo en tal dimensi�n.
                            VectorMutado(k) = Poblacion(k,j);
                        end
                    end
                    
                    %Mediante la funci�n de penalizaci�n, recalcularemos el vector mutado con anterioridad:
                    VectorMutado = Penalizacion(VectorMutado,Minimo,Maximo, Dimension);
                    %Evaluamos el vector mutado en la funci�n:
                    EvaluacionNuevoVector = NCC(Imagen,Plantilla,VectorMutado);
                    %Si esta nueva evaluaci�n es mejor que la evaluaci�n previa del individuo, entonces actualizamos el individuo por el vector mutado, incluyendo su nueva evaluaci�n.
                    if EvaluacionNuevoVector > Evaluaciones(j) 
                        Poblacion(:,j) = VectorMutado;
                        Evaluaciones(j) = EvaluacionNuevoVector;
                    end
                end
            end
        end

    end

    %   F i n    F a s e    D e    D e s a r r o l l o    D e l    A l g o r i t m o  %
    
    IndiceDelMejor = 1;
    Comparacion = 0;
    for a=1:NumeroDeIndividuos
       EvaluacionActual = Evaluaciones(a);
       if EvaluacionActual > Comparacion
          IndiceDelMejor = a;
          Comparacion = EvaluacionActual;
       end
    end

    %Actualizamos el valor de Soluci�n (el valor de la soluci�n �ptima a retornar).
    Solucion = Poblacion(:,IndiceDelMejor);

end