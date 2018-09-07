require 'nokogiri'
require 'open-uri'
require 'pry'

def parse_site(pages, uri)
  male_names = []
  female_names = []
  index = 1
  pages.times do
    mod_uri = index == 0 ? uri : "#{uri}?pg=#{index}"
    puts mod_uri
    index += 1
    site = Nokogiri::HTML(open(mod_uri))
    anchors = site.css('.entry-content').xpath('//table/tr')
    anchors.each do |a|
      el = a.css('td')
      name = el[0].text
      meaning = el[1].text
      if meaning.include?('Male')
        male_names << name
      elsif meaning.include?('Female')
        female_names << name
      elsif meaning.include?('Gender-Neutral')
        male_names << name
        female_names << name
      end
    end
  end

  { male: male_names, female: female_names }
end
