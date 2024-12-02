CREATE TABLE parcerias.users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    tipo INT NOT NULL,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);




CREATE TABLE parcerias.contratos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    cnpj VARCHAR(18) NOT NULL,
    cpf VARCHAR(14) NOT NULL,
    endereco VARCHAR(255),
    pdv VARCHAR(50),
    ativo_inativo BOOLEAN NOT NULL,
    telefone VARCHAR(20),
    whatsapp VARCHAR(20),
    email VARCHAR(255) NOT NULL,
    responsavel VARCHAR(255),
    criado_por INT NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    CONSTRAINT fk_criado_por FOREIGN KEY (criado_por)
        REFERENCES parcerias.users(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);




CREATE TABLE parcerias.empresas (
    id SERIAL PRIMARY KEY,
    descricao TEXT NOT NULL,
    razao_social VARCHAR(255) NOT NULL,
    nome_fantasia VARCHAR(255) NOT NULL,
    endereco VARCHAR(255) NOT NULL,
    id_cliente INT NOT NULL,
    fk_user_criacao INT,
    criado_em TIMESTAMP DEFAULT now(),
    CONSTRAINT fk_id_cliente FOREIGN KEY (id_cliente) 
        REFERENCES parcerias.users(id) 
        ON DELETE CASCADE,
    CONSTRAINT fk_user_criacao FOREIGN KEY (fk_user_criacao) 
        REFERENCES parcerias.users(id) 
        ON DELETE SET NULL
);




CREATE INDEX idx_empresa_id_cliente ON parcerias.empresas (id_cliente);
CREATE INDEX idx_empresa_fk_user_criacao ON parcerias.empresas (fk_user_criacao);




CREATE TABLE parcerias.filials (
  id BIGSERIAL PRIMARY KEY,
  descricao VARCHAR(255) NOT NULL,
  cnpj VARCHAR(18) UNIQUE NOT NULL,
  inscricao_estadual VARCHAR(20),
  razao_social VARCHAR(255) NOT NULL,
  nome_fantasia VARCHAR(255) NOT NULL,
  endereco VARCHAR(255) NOT NULL,
  fk_empresa BIGINT NOT NULL REFERENCES parcerias.empresas(id) ON DELETE CASCADE ON UPDATE CASCADE,
  criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  criado_por INTEGER NOT NULL,
  contribuinte BOOLEAN DEFAULT FALSE
);



CREATE TABLE parcerias.tokens_PDV (
    id SERIAL PRIMARY KEY,
    descricao VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL,
    token_pdv VARCHAR(255) UNIQUE NOT NULL,
    data_criacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    criado_por INT NOT NULL,
    FOREIGN KEY (criado_por) REFERENCES parcerias.users(id)
);




CREATE TABLE parcerias.pdvs (
    id SERIAL PRIMARY KEY,
    descricao VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL,
    data_criacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    data_ativacao_token DATE,
    data_expiracao_token DATE,
    user_criacao INT NOT NULL,
    filial_id INT NOT NULL,
    contrato_id INT NOT NULL,
    token_id INT,
    FOREIGN KEY (filial_id) REFERENCES parcerias.filials(id),
    FOREIGN KEY (contrato_id) REFERENCES parcerias.contratos(id),
    FOREIGN KEY (token_id) REFERENCES parcerias.tokens_PDV(id),
    FOREIGN KEY (user_criacao) REFERENCES parcerias.users(id)
);