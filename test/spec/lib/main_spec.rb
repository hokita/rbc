# frozen_string_literal: true

require 'main'

describe Main do
  it do
    expect(described_class.run).to eq 'test'
  end

  it do
    expect(described_class.run).to eq 'testtest'
  end
end
