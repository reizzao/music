# Estrategia_01_UM

Estrategia : {
  Nome_Estrategia: "Estrategia_01_UM",
  Inpiracao: "ZeAndre",
  Ano: 2023,
}

Sentimento_Central_escolhido: "",

Parte_A : {
  missao: "Entender o sentimentoCentral de Forma curta, forte e explicativa",
  sensacao: "Explicativa de acontecimento central",
  harmonia: "Pausada_Lenta",
  gatilho : "Frases de comando Conectar ouvinte ao sentiento central",

  frases_de_comando: {
    ACONTECEU,
    EXPLICANDO,
    RESULTADO
  }
}

Parte_B : {
  missao: "Reacao ao sentimento central",
  sensacao: "Expectativa",
  harmonia: "menor crescente",
  gatilho : "suspense/expectativa/querer sabe ro final",

  frases_de_comando: {
    DESENROLAR_REACOES,
    PORQUE,
    SOLUCAO_DESSE_DESENROLAR
  }
}

Parte_REFRAO : {
  missao: "Loop independente",
  sensacao: "Cantar levantando 1 braço só, Explosao",
  harmonia: "melodia martelo - volta sempre ao tonal",
  gatilho : "Resumao Sacada entendimento do sentimento central",

  frases_de_comando: {
    FRASE_DE_CHAMADA_OURO_FINAL,
    SACADA_ENTENDIMENTO,
    ESTADO,
    SACADA_ENTENDIMENTO_FINAL
  }
}

Parte_FECHA_LETRA : {
  missao: "moral da historia de forma envolvente",
  sensacao: "enfeite",
  harmonia: "viajante forte",
  gatilho : "moral da historia",

  frases_de_comando: {
    ENFEITE,
    SACADA_ENTENDIMENTO_MORAL_DA_HISTORIA,
    ENFEITE_2,
    SACADA_ENTENDIMENTO_MORAL_DA_HISTORIA_DO_SENTIMENTO_CENTRAL
  }
}

---
# dados:
data: 17/12/2023
categorias : [ ]

---
# harmonia_melodia
tom :
parte : A , harmonia: [ ], melodia: [ ], repete()
parte : B , harmonia: [ ], melodia: [ ], repete(2)
parte: Refrao, harmonia : [ ], melodia: [ ] , repete()
parte: Fecha_letra, harmonia: [ ], melodia: [ ] , repete()

---

