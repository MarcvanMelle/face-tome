require 'nokogiri'
require 'open-uri'
require 'pry'

def parse_site(pages, uri)
  pages += 1
  exclude = %[
    'Surnames', 'Next » ', 'All Surnames', "Celtic Surnames", "English Surnames",
    "Scottish Surnames", "Welsh Surnames", "Irish Surnames", "More lists...", "Politicians",
    "Domesday Surnames", '« Previous'
    "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"
  ]

  surnames = []
  index = 0
  pages.times do
    mod_uri = index == 0 ? uri : "#{uri}?page=#{index}"
    puts mod_uri
    index += 1
    site = Nokogiri::HTML(open(mod_uri))
    anchors = site.css('a')
    anchors.each do |anchor|
    next if anchor.attributes['href'].nil?
      if anchor.attributes['href'].value =~ /baby-names\/name-meaning/
        text = anchor.text
        next if text.to_i.positive?
        next if exclude.include?(text)
        surnames << anchor.text
      end
    end
  end
  return surnames
end

